#!/usr/bin/env ruby

skipped = [
  # disabling user functions for now
  "CreateUserFavorites",
  "GetUserFavorites",
  "GetUserInfo",
  "GetUserInfoByID",
  # wasn't working
  "GetSeasonTranslation",
]

FUNC = open("functions.go", "a+")
TEST = open("functions_test.go", "a+")

def testVar(n, t, serv)
  "var #{n} #{t} = #{testType(n, t, serv)}"
end

def testType(n, t, serv)
  return "nil" if t[0] == "*"
  return t.gsub(/^operations\./, "operations_") if t =~ /^operations\./
  return "#{serv ? "#{serv.downcase}_" : ""}#{n}_#{t}"
end

STDIN.each do |line|
  m = line.match(/func \(\w\s\*(\w+)\) (\w+)\(ctx context.Context(, )*([^\)]+)*\) \(\*operations\.(\w+), error\) \{/)
  next unless m
  next if skipped.include?(m[2])
  serv = m[1]
  serv[0] = serv[0].upcase # capitalize will lowercase the rest of the string
  params = ""
  testparams = ""
  if !m[4].nil?
    params = ", " + m[4].split(",").map { |p| p.split(" ") }.map { |p| p[0] }.join(", ")
    testparams = m[4].split(",").map { |p| p.split(" ") }.map { |p| testVar(p[0], p[1], m[1]) }.join("\t\n")
  end
  FUNC.puts <<-HERE
// #{m[2]} wraps the generated openapi.SDK#{serv != "SDK" ? ".#{serv}" : ""}.#{m[2]} call
func (c *Client) #{m[2]}(#{m[4]}) (*#{m[5]}, error) {
	r, err := c.sdk#{serv != "SDK" ? ".#{serv}" : ""}.#{m[2]}(c.ctx#{params})
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.Object, nil
}

HERE
  TEST.puts <<-HERE
func TestClient_#{m[2]}(t *testing.T) {
	c := testClient(t)
	#{testparams}
	r, err := c.#{m[2]}(#{params.split(",").drop(1).join(", ")})
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

HERE
end

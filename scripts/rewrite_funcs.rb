#!/usr/bin/env ruby

skipped = ["CreateUserFavorites"]

# lines = [
#   "func (s *artwork) GetArtworkBase(ctx context.Context, id float64) (*operations.GetArtworkBaseResponse, error) {",
#   "func (s *artwork) GetArtworkExtended(ctx context.Context, id float64) (*operations.GetArtworkExtendedResponse, error) {",
#   "func (s *artworkStatuses) GetAllArtworkStatuses(ctx context.Context) (*operations.GetAllArtworkStatusesResponse, error) {",
#   "func (s *artworkTypes) GetAllArtworkTypes(ctx context.Context) (*operations.GetAllArtworkTypesResponse, error) {",
#   "func (s *episodes) GetEpisodeExtended(ctx context.Context, id float64, meta *operations.GetEpisodeExtendedMeta) (*operations.GetEpisodeExtendedResponse, error) {",
#   "func (s *episodes) GetEpisodeTranslation(ctx context.Context, id float64, language string) (*operations.GetEpisodeTranslationResponse, error) {",
# ].each do |line|
STDIN.each do |line|
  m = line.match(/func \(\w\s\*(\w+)\) (\w+)\(ctx context.Context(, )*([^\)]+)*\) \(\*operations\.(\w+), error\) \{/)
  #puts "\n\n"+ line
  #puts m.inspect
  next unless m
  next if skipped.include?(m[2])
  serv=m[1]
  serv[0]=serv[0].upcase # capitalize will lowercase the rest of the string
  params=""
  if !m[4].nil?
    params=", " + m[4].split(',').map{|p| p.split(' ')}.map{|p| p[0]}.join(', ')
  end
  puts <<-HERE
// #{m[2]} wraps the generated openapi.SDK.#{serv}.#{m[2]} call
func (c *Client) #{m[2]}(#{m[4]}) (*#{m[5]}, error) {
	r, err := c.sdk.#{serv}.#{m[2]}(c.ctx#{params})
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, errors.Errorf("non-200 response: %d", r.StatusCode)
	}
	return r.#{m[2]}200ApplicationJSONObject, nil
}

HERE
end

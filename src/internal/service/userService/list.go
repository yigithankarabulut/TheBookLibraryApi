package userService

import "context"

func (s *userStoreService) List(ctx context.Context) (*ListUserResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		items := s.storage.List()
		response := make(ListUserResponse, len(items))

		var i int
		for _, v := range items {
			response[i] = UserResponse{
				User: v,
			}
			i++
		}
		return &response, nil
	}
}

package post

import "echo-demo-project/internal/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}

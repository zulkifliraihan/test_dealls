package requests

type SwipeActionRequest struct {
	SwipedUserID 		int `json:"swiped_user_id" binding:"required,number"`
	Direction    		string `json:"direction"  validate:"required"`
}
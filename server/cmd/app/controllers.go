package app

import "github.com/julianstephens/budget-tracker/cmd/controllers"

type ControllerSet struct {
	UserCtrl *controllers.UserCtrl
}

func newControllerSet(userCtrl *controllers.UserCtrl) *ControllerSet {
	set := &ControllerSet{
		userCtrl,
	}
	return set
}

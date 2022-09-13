package model

import (	

)

type ErrNotFound struct{

	}

func (e *ErrNotFound) Error() string {
	return "error"
}

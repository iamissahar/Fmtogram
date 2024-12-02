package tests

import "testing"

func paidPhotoReqFieldsWithName(t *testing.T)    {}
func paidPhotoReqFieldsWithoutName(t *testing.T) {}

func paidPhotoUnReqFieldsWithName(t *testing.T)    {}
func paidPhotoUnReqFieldsWithoutName(t *testing.T) {}

func paidVideoReqFieldsWithName(t *testing.T)    {}
func paidVideoReqFieldsWithoutName(t *testing.T) {}

func paidVideoUnReqFieldsWithName(t *testing.T)    {}
func paidVideoUnReqFieldsWithoutName(t *testing.T) {}

func paidmediaFunctional(t *testing.T) {
	t.Run("PaidPhotoReqFieldsWithName", paidPhotoReqFieldsWithName)
	t.Run("PaidPhotoReqFieldsWithoutName", paidPhotoReqFieldsWithoutName)
	t.Run("PaidPhotoUnReqFieldsWithName", paidPhotoUnReqFieldsWithName)
	t.Run("PaidPhotoUnReqFieldsWithoutName", paidPhotoUnReqFieldsWithoutName)
	t.Run("PaidVideoReqFieldsWithName", paidVideoReqFieldsWithName)
	t.Run("PaidVideoReqFieldsWithoutName", paidVideoReqFieldsWithoutName)
	t.Run("PaidVideoUnReqFieldsWithName", paidVideoUnReqFieldsWithName)
	t.Run("PaidVideoUnReqFieldsWithoutName", paidVideoUnReqFieldsWithoutName)
}

func TestSendPaidMedia(t *testing.T) {
	t.Run("Functional", paidmediaFunctional)
}

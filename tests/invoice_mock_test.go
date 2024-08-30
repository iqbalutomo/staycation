package test

import (
	"testing"
	"time"

	model "staycation/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type InvoiceRepoMock struct {
	mock.Mock
}

func (m *InvoiceRepoMock) CreateBooking(booking *model.Booking) error {
	args := m.Called(booking)
	return args.Error(0)
}

func (m *InvoiceRepoMock) FindBookingByID(id uint) (*model.Booking, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Booking), args.Error(1)
}

func (m *InvoiceRepoMock) FindByRoomAndDate(roomID uint, checkInDate, checkOutDate time.Time) ([]model.Booking, error) {
	args := m.Called(roomID, checkInDate, checkOutDate)
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (m *InvoiceRepoMock) CreateInvoice(invoice *model.Invoice) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func (m *InvoiceRepoMock) FindInvoiceByID(id string) (*model.Invoice, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Invoice), args.Error(1)
}

func (m *InvoiceRepoMock) UpdateInvoiceStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *InvoiceRepoMock) CreatePayment(payment *model.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func TestCreateBooking(t *testing.T) {
	repo := new(InvoiceRepoMock)
	booking := &model.Booking{
		ID: 1,
	}

	repo.On("CreateBooking", booking).Return(nil)

	err := repo.CreateBooking(booking)

	assert.Nil(t, err)
	repo.AssertExpectations(t)
}

func TestFindBookingByID(t *testing.T) {
	repo := new(InvoiceRepoMock)
	booking := &model.Booking{
		ID: 1,
	}

	repo.On("FindBookingByID", uint(1)).Return(booking, nil)

	result, err := repo.FindBookingByID(1)

	assert.Nil(t, err)
	assert.Equal(t, booking, result)
	repo.AssertExpectations(t)
}

func TestFindByRoomAndDate(t *testing.T) {
	repo := new(InvoiceRepoMock)
	checkInDate := time.Date(2024, time.August, 30, 14, 0, 0, int(time.Second), time.Local)
	checkOutDate := time.Date(2024, time.August, 31, 12, 0, 0, int(time.Second), time.Local)
	bookings := []model.Booking{
		{ID: 1},
	}

	repo.On("FindByRoomAndDate", uint(1), checkInDate, checkOutDate).Return(bookings, nil)

	result, err := repo.FindByRoomAndDate(1, checkInDate, checkOutDate)

	assert.Nil(t, err)
	assert.Equal(t, bookings, result)
	repo.AssertExpectations(t)
}

func TestCreateInvoice(t *testing.T) {
	repo := new(InvoiceRepoMock)
	invoice := &model.Invoice{
		XenditInvoiceID: "inv_123",
	}

	repo.On("CreateInvoice", invoice).Return(nil)

	err := repo.CreateInvoice(invoice)

	assert.Nil(t, err)
	repo.AssertExpectations(t)
}

func TestFindInvoiceByID(t *testing.T) {
	repo := new(InvoiceRepoMock)
	invoice := &model.Invoice{
		XenditInvoiceID: "inv_123",
	}

	repo.On("FindInvoiceByID", "inv_123").Return(invoice, nil)

	result, err := repo.FindInvoiceByID("inv_123")

	assert.Nil(t, err)
	assert.Equal(t, invoice, result)
	repo.AssertExpectations(t)
}

func TestUpdateInvoiceStatus(t *testing.T) {
	repo := new(InvoiceRepoMock)

	repo.On("UpdateInvoiceStatus", uint(1), "paid").Return(nil)

	err := repo.UpdateInvoiceStatus(1, "paid")

	assert.Nil(t, err)
	repo.AssertExpectations(t)
}

func TestCreatePayment(t *testing.T) {
	repo := new(InvoiceRepoMock)
	payment := &model.Payment{
		ID: 1,
	}

	repo.On("CreatePayment", payment).Return(nil)

	err := repo.CreatePayment(payment)

	assert.Nil(t, err)
	repo.AssertExpectations(t)
}

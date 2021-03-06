package paymentsproto

import (
	context "context"
	"time"

	"github.com/monzo/slog"
	grpc "google.golang.org/grpc"

	"swallowtail/libraries/gerrors"
)

// --- Register Payment --- //
type RegisterPaymentFuture struct {
	closer  func() error
	errc    chan error
	resultc chan *RegisterPaymentResponse
	ctx     context.Context
}

func (a *RegisterPaymentFuture) Response() (*RegisterPaymentResponse, error) {
	defer func() {
		if err := a.closer(); err != nil {
			slog.Critical(context.Background(), "Failed to close %s grpc connection: %v", "register_payment", err)
		}
	}()

	select {
	case r := <-a.resultc:
		return r, nil
	case <-a.ctx.Done():
		return nil, a.ctx.Err()
	case err := <-a.errc:
		return nil, err
	}
}

func (r *RegisterPaymentRequest) Send(ctx context.Context) *RegisterPaymentFuture {
	return r.SendWithTimeout(ctx, 10*time.Second)
}

func (r *RegisterPaymentRequest) SendWithTimeout(ctx context.Context, timeout time.Duration) *RegisterPaymentFuture {
	errc := make(chan error, 1)
	resultc := make(chan *RegisterPaymentResponse, 1)

	conn, err := grpc.DialContext(ctx, "swallowtail-s-payments:8000", grpc.WithInsecure())
	if err != nil {
		errc <- gerrors.Augment(err, "swallowtail_s_payments_connection_failed", nil)
		return &RegisterPaymentFuture{
			ctx:     ctx,
			errc:    errc,
			closer:  conn.Close,
			resultc: resultc,
		}
	}
	c := NewPaymentsClient(conn)

	ctx, cancel := context.WithTimeout(ctx, timeout)

	go func() {
		rsp, err := c.RegisterPayment(ctx, r)
		if err != nil {
			errc <- gerrors.Augment(err, "failed_register_payment", nil)
			return
		}
		resultc <- rsp
	}()

	return &RegisterPaymentFuture{
		ctx: ctx,
		closer: func() error {
			cancel()
			return conn.Close()
		},
		errc:    errc,
		resultc: resultc,
	}
}

// --- Read Users Last Payment  --- //

type ReadUsersLastPaymentFuture struct {
	closer  func() error
	errc    chan error
	resultc chan *ReadUsersLastPaymentResponse
	ctx     context.Context
}

func (a *ReadUsersLastPaymentFuture) Response() (*ReadUsersLastPaymentResponse, error) {
	defer func() {
		if err := a.closer(); err != nil {
			slog.Critical(context.Background(), "Failed to close %s grpc connection: %v", "read_users_last_payment", err)
		}
	}()

	select {
	case r := <-a.resultc:
		return r, nil
	case <-a.ctx.Done():
		return nil, a.ctx.Err()
	case err := <-a.errc:
		return nil, err
	}
}

func (r *ReadUsersLastPaymentRequest) Send(ctx context.Context) *ReadUsersLastPaymentFuture {
	return r.SendWithTimeout(ctx, 10*time.Second)
}

func (r *ReadUsersLastPaymentRequest) SendWithTimeout(ctx context.Context, timeout time.Duration) *ReadUsersLastPaymentFuture {
	errc := make(chan error, 1)
	resultc := make(chan *ReadUsersLastPaymentResponse, 1)

	conn, err := grpc.DialContext(ctx, "swallowtail-s-payments:8000", grpc.WithInsecure())
	if err != nil {
		errc <- gerrors.Augment(err, "swallowtail_s_payments_connection_failed", nil)
		return &ReadUsersLastPaymentFuture{
			ctx:     ctx,
			errc:    errc,
			closer:  conn.Close,
			resultc: resultc,
		}
	}
	c := NewPaymentsClient(conn)

	ctx, cancel := context.WithTimeout(ctx, timeout)

	go func() {
		rsp, err := c.ReadUsersLastPayment(ctx, r)
		if err != nil {
			errc <- gerrors.Augment(err, "failed_read_users_last_payment", nil)
			return
		}
		resultc <- rsp
	}()

	return &ReadUsersLastPaymentFuture{
		ctx: ctx,
		closer: func() error {
			cancel()
			return conn.Close()
		},
		errc:    errc,
		resultc: resultc,
	}
}

// --- List Payments By User ID --- //

type ListPaymentsByUserIDFuture struct {
	closer  func() error
	errc    chan error
	resultc chan *ListPaymentsByUserIDResponse
	ctx     context.Context
}

func (a *ListPaymentsByUserIDFuture) Response() (*ListPaymentsByUserIDResponse, error) {
	defer func() {
		if err := a.closer(); err != nil {
			slog.Critical(context.Background(), "Failed to close %s grpc connection: %v", "list_payments_by_user_id", err)
		}
	}()

	select {
	case r := <-a.resultc:
		return r, nil
	case <-a.ctx.Done():
		return nil, a.ctx.Err()
	case err := <-a.errc:
		return nil, err
	}
}

func (r *ListPaymentsByUserIDRequest) Send(ctx context.Context) *ListPaymentsByUserIDFuture {
	return r.SendWithTimeout(ctx, 10*time.Second)
}

func (r *ListPaymentsByUserIDRequest) SendWithTimeout(ctx context.Context, timeout time.Duration) *ListPaymentsByUserIDFuture {
	errc := make(chan error, 1)
	resultc := make(chan *ListPaymentsByUserIDResponse, 1)

	conn, err := grpc.DialContext(ctx, "swallowtail-s-payments:8000", grpc.WithInsecure())
	if err != nil {
		errc <- gerrors.Augment(err, "swallowtail_s_payments_connection_failed", nil)
		return &ListPaymentsByUserIDFuture{
			ctx:     ctx,
			errc:    errc,
			closer:  conn.Close,
			resultc: resultc,
		}
	}
	c := NewPaymentsClient(conn)

	ctx, cancel := context.WithTimeout(ctx, timeout)

	go func() {
		rsp, err := c.ListPaymentsByUserID(ctx, r)
		if err != nil {
			errc <- gerrors.Augment(err, "failed_list_payments_by_user_id", nil)
			return
		}
		resultc <- rsp
	}()

	return &ListPaymentsByUserIDFuture{
		ctx: ctx,
		closer: func() error {
			cancel()
			return conn.Close()
		},
		errc:    errc,
		resultc: resultc,
	}
}

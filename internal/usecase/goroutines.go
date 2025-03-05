package usecase

/*
func (u *Usecase) UpdatePayments() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// p.s по хорошему здесь делаем запрос в бд и получаем payment_status_id где payment_status=cancel
	newPaymentStatusID := 3
	timeNow := time.Now()
	for {
		select {
		case <-ticker.C:
			err := u.pgRepo.UpdatePayments(ctx, newPaymentStatusID, timeNow)
			if err != nil {
				u.logger.Error("usecase UpdatePayments u.pgRepo.UpdatePayments", slog.Any("error", err))
				return
			}
		case <-ctx.Done():
			u.logger.Info("UpdatePayments stopped")
			return
		}
	}
}

func (u *Usecase) UpdateSubscriptions() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	newSubscriptionStatusID := 1
	timeNow := time.Now()
	defer cancel()
	// p.s по хорошему здесь делаем запрос в бд и получаем payment_status_id где subscription_status=0
	for {
		select {
		case <-ticker.C:
			err := u.pgRepo.UpdateSubscriptions(ctx, newSubscriptionStatusID, timeNow)
			if err != nil {
				u.logger.Error("usecase UpdateSubscriptions u.pgRepo.UpdateSubscriptions", slog.Any("error", err))
				return
			}
		case <-ctx.Done():
			u.logger.Info("UpdateSubscriptions stopped")
			return
		}
	}
}
*/

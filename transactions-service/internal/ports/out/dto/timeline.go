package dto

import "transactions-service/internal/core/domain"

type TimelineResponse []*TransactionResponse

func TimelineResponseFromTransactionList(transactions []*domain.Transaction) TimelineResponse {
	timeline := make(TimelineResponse, 0)
	for _, transaction := range transactions {
		timeline = append(timeline, TransactionResponseFromDomain(transaction))
	}
	return timeline
}

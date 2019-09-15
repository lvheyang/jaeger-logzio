package store

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"time"
)

type logzioSpanReader struct {
	apiToken string
	logger   hclog.Logger
}

func NewLogzioSpanReader(config LogzioConfig, logger hclog.Logger) *logzioSpanReader {
	return &logzioSpanReader{
		logger:   logger,
		apiToken: config.Api_Token,
	}
}

func (sr *logzioSpanReader) GetTrace(ctx context.Context, traceID model.TraceID) (*model.Trace, error) {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "GetTrace")
	//defer span.Finish()
	//currentTime := time.Now()
	//traces, err := sr.multiRead(ctx, []model.TraceID{traceID}, currentTime.Add(time.Hour*48), currentTime)
	//if err != nil {
	//	return nil, err
	//}
	//if len(traces) == 0 {
	//	return nil, spanstore.ErrTraceNotFound
	//}
	//return traces[0], nil
	return nil, nil
}

func (*logzioSpanReader) GetServices(ctx context.Context) ([]string, error) {
	return nil, nil
}

func (*logzioSpanReader) GetOperations(ctx context.Context, service string) ([]string, error) {
	return nil, nil
}

func (*logzioSpanReader) FindTraces(ctx context.Context, query *spanstore.TraceQueryParameters) ([]*model.Trace, error) {
	return nil, nil
}

func (*logzioSpanReader) FindTraceIDs(ctx context.Context, query *spanstore.TraceQueryParameters) ([]model.TraceID, error) {
	return nil, nil
}

func (*logzioSpanReader) GetDependencies(endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error) {
	return nil, nil
}

//func (sr *logzioSpanReader) multiRead(ctx context.Context, traceIDs []model.TraceID, startTime, endTime time.Time) ([]*model.Trace, error) {
//
//	childSpan, _ := opentracing.StartSpanFromContext(ctx, "multiRead")
//	childSpan.LogFields(otlog.Object("trace_ids", traceIDs))
//	defer childSpan.Finish()
//
//	if len(traceIDs) == 0 {
//		return []*model.Trace{}, nil
//	}
//
//	// Add an hour in both directions so that traces that straddle two indexes are retrieved.
//	// i.e starts in one and ends in another.
//	nextTime := model.TimeAsEpochMicroseconds(startTime.Add(-time.Hour))
//
//	searchAfterTime := make(map[model.TraceID]uint64)
//	totalDocumentsFetched := make(map[model.TraceID]int)
//	tracesMap := make(map[model.TraceID]*model.Trace)
//	for {
//		if len(traceIDs) == 0 {
//			break
//		}
//		searchRequests := make([]*elastic.SearchRequest, len(traceIDs))
//		for i, traceID := range traceIDs {
//			query := elastic.NewTermQuery("traceID", traceID.String())
//			if val, ok := searchAfterTime[traceID]; ok {
//				nextTime = val
//			}
//
//			s := sr.sourceFn(query, nextTime)
//
//			searchRequests[i] = elastic.NewSearchRequest().
//				IgnoreUnavailable(true).
//				Type(spanType).
//				Source(s)
//		}
//		// set traceIDs to empty
//		traceIDs = nil
//		results, err := sr.client.MultiSearch().Add(searchRequests...).Index(indices...).Do(sr.ctx)
//
//		if err != nil {
//			logErrorToSpan(childSpan, err)
//			return nil, err
//		}
//
//		if results.Responses == nil || len(results.Responses) == 0 {
//			break
//		}
//
//		for _, result := range results.Responses {
//			if result.Hits == nil || len(result.Hits.Hits) == 0 {
//				continue
//			}
//			spans, err := sr.collectSpans(result.Hits.Hits)
//			if err != nil {
//				logErrorToSpan(childSpan, err)
//				return nil, err
//			}
//			lastSpan := spans[len(spans)-1]
//
//			if traceSpan, ok := tracesMap[lastSpan.TraceID]; ok {
//				traceSpan.Spans = append(traceSpan.Spans, spans...)
//			} else {
//				tracesMap[lastSpan.TraceID] = &model.Trace{Spans: spans}
//			}
//
//			totalDocumentsFetched[lastSpan.TraceID] = totalDocumentsFetched[lastSpan.TraceID] + len(result.Hits.Hits)
//			if totalDocumentsFetched[lastSpan.TraceID] < int(result.TotalHits()) {
//				traceIDs = append(traceIDs, lastSpan.TraceID)
//				searchAfterTime[lastSpan.TraceID] = model.TimeAsEpochMicroseconds(lastSpan.StartTime)
//			}
//		}
//	}
//
//	var traces []*model.Trace
//	for _, trace := range tracesMap {
//		traces = append(traces, trace)
//	}
//	return traces, nil
//}
package stats

// Publisher
type Publisher interface {
    Publish(name string, metric int64)
}

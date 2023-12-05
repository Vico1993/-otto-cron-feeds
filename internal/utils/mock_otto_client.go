package utils

// func SetupClient() otto.FeedService {
// 	mockFeedService := new(MockFeedClient)
// 	return mockFeedService
// }

// type MockFeedClient struct {
// 	mock.Mock
// 	client http.Client
// }

// func (m *MockFeedClient) Create(url string) *otto.Feed {
// 	args := m.Called(url)
// 	return args.Get(0).(*otto.Feed)
// }

// func (m *MockFeedClient) Link(chatId string, threadId string, feedId string) bool {
// 	args := m.Called(chatId, threadId, feedId)
// 	return args.Get(0).(bool)
// }

// func (m *MockFeedClient) List(chatId string, threadId string) []otto.Feed {
// 	args := m.Called(chatId, threadId)
// 	return args.Get(0).([]otto.Feed)
// }

// func (m *MockFeedClient) ListAll(active bool) []otto.Feed {
// 	args := m.Called(active)
// 	return args.Get(0).([]otto.Feed)
// }

// func (m *MockFeedClient) ListArticles(feedId string) []otto.Article {
// 	args := m.Called(feedId)
// 	return args.Get(0).([]otto.Article)
// }

// func (m *MockFeedClient) UnLink(chatId string, threadId string, feedId string) bool {
// 	args := m.Called(chatId, threadId, feedId)
// 	return args.Get(0).(bool)
// }

package download

type Document struct {
	url  string
	html string
}

func NewDocument(url, html string) *Document {
	return &Document{
		url:  url,
		html: html,
	}
}

/*
第一，构造函数中如果初始化downloader.downloadHtml()逻辑复杂，耗时⻓，不应该放到构造函数中，会影响代码的可测试性。
第二，如果HtmlDownloader对象在构造函数中 通过new来创建，违反了基于接口而非实现编程的设计思想，也会影响到代码的可测试性。
第三，从业务含 义上来讲，Document网⻚文档没必要依赖HtmlDownloader类，违背了迪米特法则。
*/

type DocumentFactory struct {
	downloader HtmlDownloader
}

func NewDocumentFactory(downloader HtmlDownloader) *DocumentFactory {
	return &DocumentFactory{
		downloader: downloader,
	}
}

func (self *DocumentFactory) CreateDocument(url string) (*Document, error) {
	html, err := self.downloader.downloadHtml(url)
	if err != nil {
		return nil, err
	}
	return NewDocument(url, html), nil
}

# Ginについて

## 環境構築

[Quickstart](https://gin-gonic.com/docs/quickstart/)

## アーキテクチャ
- Controller
    リクエストデータのハンドリングやレスポンスの設定
- Service
    実現したい機能(ビジネスロジック)の実装
- Repository
    データの永続化やデータソースとのやり取り

Router -> IController <- Controller -> IService <- Service -> IRepsitory <- Repository -> Data
IController, IService, IRepositoryはインターフェースで実装することで依存関係が小さくなる

## SOLIDの依存関係の逆転の原則で実装する

```go
// Router


// RepositoryImpl -> Data
type ItemRepository struct {
	items []models.Item
}

// Find関数を書くだけでInterfaceを実装したことになる
func (i *ItemRepository) Find() (*[]models.Item, error){
    return i.items,nil
}

// 公開関数
func NewItemRepository(items []models.Item) IItemRepository {
    return &ItemRepository{items: items}
}

// Repository Interface <-  Repository
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

// IRepository <- Service で紐づける
type ItemService struct {
	repositories repositories.IItemRepository
}

// Controller -> IItemService <- Service
type IItemService interface {
    FindAll() (*[]models.Item, error)
}
// Controller...

func (i *ItemService) FindAll() (*[]models.Item, error){
    return repository.Find() //return (*[]models.Item, error)
}

```

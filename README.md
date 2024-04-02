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


## dockerの追加
- postgres環境の構築
- pgadmin(DB管理画面)の構築

```shell
$ docker compose up -d
```

## ORMを使う
- ORMを学習する
- パフォーマンスチューニングが難しい

## DB Migration
- カラムの変更
- インデックスの適用
- ロールバック機能がある
- 開発段階でカラムの変更がある

## DBの設定
- `.env`ファイルをローカルに作成する

```
DB_HOST=localhost
DB_USER=xxx
DB_PASSWORD=xxx
DB_NAME=xxx
DB_PORT=5432
```

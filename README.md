# Ginについて

## 目次
- [Ginについて](#ginについて)
  - [目次](#目次)
  - [Memo](#memo)
  - [環境構築](#環境構築)
  - [アーキテクチャ](#アーキテクチャ)
  - [SOLIDの依存関係の逆転の原則で実装する](#solidの依存関係の逆転の原則で実装する)
  - [dockerの追加](#dockerの追加)
  - [ORMを使う](#ormを使う)
  - [DB Migration](#db-migration)
  - [DBの設定](#dbの設定)
  - [ルーティングのグループ化](#ルーティングのグループ化)
  - [JWTの概要](#jwtの概要)
    - [構成](#構成)
    - [認証](#認証)
  - [備考, Memo](#備考-memo)
  - [jwt認証を導入](#jwt認証を導入)
  - [ミドルウェア](#ミドルウェア)
  - [CORSについて](#corsについて)
    - [設定](#設定)
    - [参考](#参考)
  - [GraphQLを導入してみた](#graphqlを導入してみた)
  - [ディレクトリ構成](#ディレクトリ構成)

## Memo
- `README.md`
- `ginMemo/`


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

## ルーティングのグループ化
```go
g := gin.Default()
itemRouter:=g.Group("/items")
itemRouter.Get("",gin.HandlerFunc)
authRouter:=g.Group("/auth")
```

## JWTの概要

### 構成
- head
    ハッシュアルゴリズムの情報などのメタデータ
- payload
    認証対象の情報でユーザー名やIDなどの任意の情報
- signature
    ヘッダとペイロードをエンコードしたものに秘密鍵を加えてハッシュ化した値
    head + payload + (Encode_head + Encode_payload + pr_key) --> ハッシュ化

### 認証
![Jwt](<jwt_image.png>)

## 備考, Memo

- DTOはRequestデータを`model/**`に変換するために用いる
```go
	var input dto.SignupInput
    //Request DataとDTO.SignInput型をbindingして、Validationも行う
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
```

- Interfaceを分離させておくことで実装を細かく分割できる

```go

type ISignup interface {
	Signup(email string, password string) error
}
type ISignout interface {
	Signout(email string, password string) error
}
type GoogleAuthenticate interface {
	ISignup
	ISignout
}

type FacebookAuthenticate interface {
	ISignup
	ISignout
}

```
## jwt認証を導入


## ミドルウェア

リクエストの途中で処理を挟むことができる

`Gin`のミドルウェアをエンドポイントごとに設定できる
1. 全部のエンドポイント
2. 個別のエンドポイント
3. グルーピング(Group)したエンドポイント

```go
g := Gin.Default()
// 1
g.Use(Middleware)
// 2
g.GET("/",Middleware,HandlerFunc)
// 3
GroupRouter.Use(Middleware)

```

## CORSについて

オリジンとは、`protocol` + `domain` + `port number`
- 異なるオリジン間のデータリソースのやり取りを制限するもの
- フロント、バックエンドサーバーを分けているため、サーバーごとでオリジンが制限をしている
  なので、オリジン間のリソースの共有をする
> あるオリジンで動いている Web アプリケーションに対して、別のオリジンのサーバーへのアクセスをオリジン間 HTTP リクエストによって許可できる仕組み

### 設定

```go
import (
  "time"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)


func main() {
  router := gin.Default()
  // CORS for https://foo.com and https://github.com origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
  // - Credentials share
  // - Preflight requests cached for 12 hours
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://foo.com"},
    AllowMethods:     []string{"PUT", "PATCH","GET","POST"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))
  router.Run()
}
```

### 参考
- [なんとなく CORS がわかる...はもう終わりにする]("https://qiita.com/att55/items/2154a8aad8bf1409db2b")

- [GIN CORS Middleware (公式リポジトリ)](https://github.com/gin-contrib/cors?tab=readme-ov-file#cors-gins-middleware)


## GraphQLを導入してみた
[graphql.md](./graphql.md)に記述


## ディレクトリ構成

```shell
.
├── GinMemo
│   └── sub.go
├── README.md
├── docker-compose.yaml
├── go.mod
├── go.sum
├── gqlgen.yaml
├── graph_example
│   ├── graph
│   │   ├── model
│   │   │   └── models_gen.go
│   │   └── resolver.go
│   ├── internal
│   │   └── generated.go
│   └── schema.graphqls
├── graphql.md
├── jwt_image.png
├── main.go
└── mock
    ├── controllers
    │   ├── auth_controller.go
    │   └── item_controller.go
    ├── dto
    │   ├── auth_dto.go
    │   └── item_dto.go
    ├── infra
    │   ├── db.go
    │   └── initializer.go
    ├── item
    ├── migrations
    │   └── migration.go
    ├── models
    │   ├── item.go
    │   └── user.go
    ├── repositories
    │   ├── auth_repository.go
    │   └── item_repository.go
    └── services
        ├── auth_service.go
        └── item_service.go

```

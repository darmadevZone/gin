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

Router -> IController <- Controller -> IService <- Service -> IRepository <- Repository -> Data
IController, IService, IRepositoryはインターフェースで実装することで依存関係が小さくなる

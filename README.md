# AWS cost notification

AWSのコストを月1でLINEに通知するLambda関数。

## 環境構築

```sh
cp sample.env.json env.json
```

```json
{
  "LINE_ACCESS_TOKEN": "LINEのアクセストークン",
  "OPEN_EXCHANGE_RATES_APP_ID": "Open Exchange RatesのアプリケーションID"
}
```

## ローカル実行

```make build```でビルド後に以下のコマンドを実行する。

```sh
make local_invoke
```

## デプロイ

初回デプロイ時に以下のコマンドを実行して、```samconfig.toml```を作成する。

```sh
sam deploy --parameter-overrides LineAccessToken=$LINE_ACCESS_TOKEN OpenExchangeRatesAppId=$OPEN_EXCHANGE_RATES_APP_ID --guided
```

デプロイを実行する。

```sh
make deploy
```

**samconfig.tomlについて**

AWS SAM CLIの設定ファイルであり、samコマンドを使用する際の設定が定義されたファイル。
sam deploy --guided コマンドを使用した際にプロジェクトのルートディレクトリに自動生成される。

デプロイに関する設定が記述されており、```parameter_overrides```にデプロイ時に渡すパラメータが記述することで、環境変数を渡すことができる。

# Phase 1: MVPコア機能
**目的**: 求人一覧・詳細表示、企業掲載依頼受付、管理画面での手動審査、DB設計など基本部分を構築。

**バックエンドタスク**

- [BE-1] DBスキーマ設計
    - テーブル例:
        - `companies`（企業情報, 認証状態フラグ含む）
        - `jobs`（求人情報: 雇用形態タグ1つまで, 報酬、技術スタックなど）
        - `users`（ログインユーザ用, OAuth情報・独自SignUp対応）
        - `inquiries`（掲載依頼フォームのログ）
        - `auth_tokens`（JWTやOAuthのトークン格納, 必要なら）
    - `jobs`テーブルで雇用形態フィールドはENUMか、typeカラム（`intern`, `fulltime`など）を固定値とする。
- [BE-2] Goエントリーポイント実装・APIベースセットアップ
    - ディレクトリ構成: `cmd/`, `internal/`, `pkg/`等
    - ルーティング: chiやEchoでルート設定
- [BE-3] 求人一覧取得API (GET /api/jobs)
    - クエリパラメータで検索(技術名、キーワード)、絞り込み(雇用形態タグ)に対応
    - JSONでレスポンス
- [BE-4] 求人詳細取得API (GET /api/jobs/{id})
    - 該当求人の詳細情報返す
- [BE-5] 企業掲載依頼(問い合わせ)API (POST /api/inquiries)
    - 未認証企業が掲載依頼するためのフォーム受付
    - データ保存とSlack通知 or メール送信
- [BE-6] 管理画面用API (GET /api/admin/inquiries, PUT /api/admin/inquiries/{id}/approve)
    - 管理者がinquiryを見て、企業を`companies`テーブルへ承認追加
    - 企業承認後は`companies`の`is_verified`フラグをtrueに

**フロントエンド(最低限, htmx利用)**

- [FE-1] `index.html`作成済み（Landingページ）
    - Jobs一覧表示部分 `/featured_companies` などMockをAPI化
    - Testimonialsも同様
- [FE-2] Jobs検索UI(hx-getで/api/jobsを叩いて一覧表示)
- [FE-3] Job詳細表示UI(hx-getで/api/jobs/{id}叩いてモーダルや別ページ表示)
- [FE-4] InquiryフォームUI(hx-postで/api/inquiriesに送信)

**管理UI(最低限HTML + htmx)**

- [FE-5] /adminページで承認待ち企業一覧(hx-get: /api/admin/inquiries)
- [FE-6] /adminページでapproveボタン押下でPUT /api/admin/inquiries/{id}/approve 呼び出し

**その他**

- [DEV-1] CI/CDセットアップ(GitHub Actions)
    - Go test & Lint実行
    - Docker image build & push
- [DEV-2] Docker Compose or Terraformで環境構築
    - Goバイナリ + PostgreSQL（or MySQL）
- [DEV-3] 簡易的なエラーハンドリング・ログ出力設定(logrus, zapなど利用)


# Phase 2: 認証・認可、企業ログインフロー・認証済み企業の求人投稿

**目的**: 企業向けに認証機能追加、認証済み企業が管理画面から直接求人を投稿できるようにする。

**バックエンドタスク**

- [BE-7] OAuth連携(Google, GitHub, Apple)によるUserログインAPI実装
- [BE-8] 独自SignUp/Login API実装(必要なら)
- [BE-9] 認証済み企業が /api/companies/me/jobs (POST)で求人登録するAPI実装
    - 認証済みの場合、審査なしで直ちに`jobs`テーブルにINSERT

**フロントタスク**

- [FE-7] ログインUI（Sign inボタンでOAuth画面へリダイレクト）
- [FE-8] 企業ログイン後の専用UI(/mycompany)で求人投稿フォーム作成(hx-post: /api/companies/me/jobs)


# Phase 3: さらなる改善(オプション)

**目的**: サブスクモデル、求人上位表示などビジネス拡張、ユーザーの過去閲覧キャッシュ機能。

**バックエンドタスク**

- [BE-10] 有料サブスクプランAPI(Stripe等で決済連携)
- [BE-11] 求人上位表示のパラメータ追加(閲覧時のソート順でプレミア企業が上位に)
- [BE-12] ユーザー閲覧履歴記録API（ログイン済みの場合）

**フロントタスク**

- [FE-9] プラン選択UI(企業向け)
- [FE-10] ユーザーダッシュボードで過去の閲覧求人一覧(hx-get: /api/users/me/history)

---

### 収益モデル追加アイデア

- 基本掲載料(シンプルな1求人あたり定額)
- 有料サブスク：月額で複数求人を同時掲載可能 + 上位表示オプション
- 控えめな広告枠：業界関連サービスツールベンダーの小さなロゴ広告など
- 紹介料モデル：マッチング成功時にフィーをいただく(後から実装)

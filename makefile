up: # イメージを起動
	docker-compose up -d
stop: # イメージを止める
	docker-compose stop
destroy: # すべてのボリュームを削除する
	docker-compose down --rmi all --volumes
ps: # 動いているコンテナの確認
	docker-compose ps
front: # サーバ用のコンテナにアクセスする
	docker-compose exec client bash
app: # サーバ用のコンテナにアクセスする
	docker-compose exec server bash
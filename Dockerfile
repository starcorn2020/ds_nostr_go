# 使用官方 PostgreSQL 鏡像作為基礎映像
FROM postgres:latest

# 設定環境變量
ENV POSTGRES_PASSWORD=0000

# 將本地目錄掛載到容器中
VOLUME /var/lib/postgresql/data

# 目錄的權限設定為容器運行時的用戶和組
RUN chown -R postgres:postgres /var/lib/postgresql/data

# 設定容器的用戶和組
USER postgres:postgres

# 運行初始化指令
RUN initdb --locale en_US.utf8 -E UTF8 -D /var/lib/postgresql/data

# 指定容器運行時要執行的命令
CMD ["postgres"]

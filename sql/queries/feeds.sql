-- name: CreateFeed :one 
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
select * from feeds;

-- name: GetNextFeedToFetch :one
select * from feeds
order by list_fetched_at asc nulls first
limit 1;

-- name: MarkFeedAsFetched :one
update feeds set list_fetched_at = now(),
updated_at = now()
where id = $1
RETURNING *;
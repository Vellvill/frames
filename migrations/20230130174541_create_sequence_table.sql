-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sequence
(
    good_id BIGINT,
    image_id BIGINT,
    sequence INT,

    CONSTRAINT fk_image_id
        FOREIGN KEY(image_id)
            REFERENCES public.image(id),
    CONSTRAINT fk_good_id
        FOREIGN KEY(good_id)
            REFERENCES public.good(id)
);

create unique index on sequence(image_id, sequence);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sequence;
-- +goose StatementEnd

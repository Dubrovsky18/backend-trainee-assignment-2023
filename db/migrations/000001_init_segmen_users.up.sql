CREATE TABLE "Slug"
(
    "name_slug" VARCHAR(255),
);

CREATE TABLE "Users"
(
    "id" BIGSERIAL generated always as identity PRIMARY KEY,
    "slugs" []VARCHAR(255) references "Slug",
);


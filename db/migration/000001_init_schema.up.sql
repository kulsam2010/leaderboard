CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "scores" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "activity_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "score_audit" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "activity_id" integer NOT NULL,
  "points" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "activities" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "scores" ("user_id");

CREATE INDEX ON "scores" ("user_id", "activity_id");

CREATE INDEX ON "score_audit" ("user_id");

CREATE INDEX ON "score_audit" ("activity_id");

CREATE INDEX ON "score_audit" ("user_id", "activity_id");

ALTER TABLE "scores" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "scores" ADD FOREIGN KEY ("activity_id") REFERENCES "activities" ("id");

ALTER TABLE "score_audit" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "score_audit" ADD FOREIGN KEY ("activity_id") REFERENCES "activities" ("id");

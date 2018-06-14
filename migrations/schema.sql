--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE authors (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    photo_url character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE authors OWNER TO postgres;

--
-- Name: authors_snacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE authors_snacks (
    id uuid NOT NULL,
    snack_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE authors_snacks OWNER TO postgres;

--
-- Name: episodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE episodes (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    markdown character varying(255) NOT NULL,
    thumbnail_url character varying(255) NOT NULL,
    embed_code character varying(255) NOT NULL,
    body character varying(255) NOT NULL,
    pro boolean NOT NULL,
    repo character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE episodes OWNER TO postgres;

--
-- Name: episodes_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE episodes_authors (
    id uuid NOT NULL,
    episode_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE episodes_authors OWNER TO postgres;

--
-- Name: episodes_series; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE episodes_series (
    id uuid NOT NULL,
    episode_id uuid NOT NULL,
    series_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE episodes_series OWNER TO postgres;

--
-- Name: episodes_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE episodes_topics (
    id uuid NOT NULL,
    episode_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE episodes_topics OWNER TO postgres;

--
-- Name: guides; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE guides (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    markdown character varying(255) NOT NULL,
    thumbnail_url character varying(255) NOT NULL,
    embed_code character varying(255) NOT NULL,
    body character varying(255) NOT NULL,
    pro boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE guides OWNER TO postgres;

--
-- Name: guides_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE guides_authors (
    id uuid NOT NULL,
    guide_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE guides_authors OWNER TO postgres;

--
-- Name: guides_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE guides_topics (
    id uuid NOT NULL,
    guide_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE guides_topics OWNER TO postgres;

--
-- Name: images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE images (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    alt_text character varying(255) NOT NULL,
    file_name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE images OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: series; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE series (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    thumbnail_url character varying(255) NOT NULL,
    body character varying(255) NOT NULL,
    pro boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE series OWNER TO postgres;

--
-- Name: series_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE series_authors (
    id uuid NOT NULL,
    series_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE series_authors OWNER TO postgres;

--
-- Name: series_episodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE series_episodes (
    id uuid NOT NULL,
    series_id uuid NOT NULL,
    episode_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE series_episodes OWNER TO postgres;

--
-- Name: series_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE series_topics (
    id uuid NOT NULL,
    series_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE series_topics OWNER TO postgres;

--
-- Name: snacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE snacks (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    sponsored boolean NOT NULL,
    url character varying(255) NOT NULL,
    summary character varying(255) NOT NULL,
    comment character varying(255) NOT NULL,
    embed_code character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE snacks OWNER TO postgres;

--
-- Name: topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE topics (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE topics OWNER TO postgres;

--
-- Name: topics_snacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE topics_snacks (
    id uuid NOT NULL,
    snack_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE topics_snacks OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    admin boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- Name: authors_snacks authors_snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY authors_snacks
    ADD CONSTRAINT authors_snacks_pkey PRIMARY KEY (id);


--
-- Name: episodes_authors episodes_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY episodes_authors
    ADD CONSTRAINT episodes_authors_pkey PRIMARY KEY (id);


--
-- Name: episodes episodes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY episodes
    ADD CONSTRAINT episodes_pkey PRIMARY KEY (id);


--
-- Name: episodes_series episodes_series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY episodes_series
    ADD CONSTRAINT episodes_series_pkey PRIMARY KEY (id);


--
-- Name: episodes_topics episodes_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY episodes_topics
    ADD CONSTRAINT episodes_topics_pkey PRIMARY KEY (id);


--
-- Name: guides_authors guides_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY guides_authors
    ADD CONSTRAINT guides_authors_pkey PRIMARY KEY (id);


--
-- Name: guides guides_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY guides
    ADD CONSTRAINT guides_pkey PRIMARY KEY (id);


--
-- Name: guides_topics guides_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY guides_topics
    ADD CONSTRAINT guides_topics_pkey PRIMARY KEY (id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: series_authors series_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY series_authors
    ADD CONSTRAINT series_authors_pkey PRIMARY KEY (id);


--
-- Name: series_episodes series_episodes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY series_episodes
    ADD CONSTRAINT series_episodes_pkey PRIMARY KEY (id);


--
-- Name: series series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY series
    ADD CONSTRAINT series_pkey PRIMARY KEY (id);


--
-- Name: series_topics series_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY series_topics
    ADD CONSTRAINT series_topics_pkey PRIMARY KEY (id);


--
-- Name: snacks snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY snacks
    ADD CONSTRAINT snacks_pkey PRIMARY KEY (id);


--
-- Name: topics topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (id);


--
-- Name: topics_snacks topics_snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY topics_snacks
    ADD CONSTRAINT topics_snacks_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--


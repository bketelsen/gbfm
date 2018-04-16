--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
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


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    photo_url character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors OWNER TO postgres;

--
-- Name: episodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.episodes (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    thumbnail_url character varying(255) NOT NULL,
    embed_code text NOT NULL,
    body character varying(255) NOT NULL,
    pro boolean NOT NULL,
    repo character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    display_order integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.episodes OWNER TO postgres;

--
-- Name: episodes_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.episodes_authors (
    id uuid NOT NULL,
    episode_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.episodes_authors OWNER TO postgres;

--
-- Name: episodes_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.episodes_topics (
    id uuid NOT NULL,
    episode_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.episodes_topics OWNER TO postgres;

--
-- Name: gbfm; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gbfm (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    embed_code character varying(255) NOT NULL,
    github_link character varying(255) NOT NULL,
    sponsor character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.gbfm OWNER TO postgres;

--
-- Name: gifm_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gifm_authors (
    id uuid NOT NULL,
    topic_id uuid NOT NULL,
    gifm_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.gifm_authors OWNER TO postgres;

--
-- Name: gifm_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gifm_topics (
    id uuid NOT NULL,
    topic_id uuid NOT NULL,
    gifm_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.gifm_topics OWNER TO postgres;

--
-- Name: guides; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.guides (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    thumbnail_url character varying(255) NOT NULL,
    embed_code character varying(255) NOT NULL,
    body character varying(255) NOT NULL,
    pro boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.guides OWNER TO postgres;

--
-- Name: guides_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.guides_authors (
    id integer NOT NULL,
    guide_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.guides_authors OWNER TO postgres;

--
-- Name: guides_authors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.guides_authors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.guides_authors_id_seq OWNER TO postgres;

--
-- Name: guides_authors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.guides_authors_id_seq OWNED BY public.guides_authors.id;


--
-- Name: guides_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.guides_topics (
    id integer NOT NULL,
    guide_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.guides_topics OWNER TO postgres;

--
-- Name: guides_topics_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.guides_topics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.guides_topics_id_seq OWNER TO postgres;

--
-- Name: guides_topics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.guides_topics_id_seq OWNED BY public.guides_topics.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: series; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series (
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


ALTER TABLE public.series OWNER TO postgres;

--
-- Name: series_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series_authors (
    id uuid NOT NULL,
    series_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.series_authors OWNER TO postgres;

--
-- Name: series_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.series_topics (
    id uuid NOT NULL,
    series_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.series_topics OWNER TO postgres;

--
-- Name: snacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.snacks (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    sponsored boolean NOT NULL,
    url character varying(255) NOT NULL,
    embed_code text NOT NULL,
    summary character varying(255) NOT NULL,
    comment character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.snacks OWNER TO postgres;

--
-- Name: snacks_authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.snacks_authors (
    id uuid NOT NULL,
    snack_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.snacks_authors OWNER TO postgres;

--
-- Name: snacks_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.snacks_topics (
    id uuid NOT NULL,
    snack_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.snacks_topics OWNER TO postgres;

--
-- Name: topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics (
    id uuid NOT NULL,
    slug character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.topics OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    admin boolean DEFAULT false NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: guides_authors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.guides_authors ALTER COLUMN id SET DEFAULT nextval('public.guides_authors_id_seq'::regclass);


--
-- Name: guides_topics id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.guides_topics ALTER COLUMN id SET DEFAULT nextval('public.guides_topics_id_seq'::regclass);


--
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- Name: episodes_authors episodes_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.episodes_authors
    ADD CONSTRAINT episodes_authors_pkey PRIMARY KEY (id);


--
-- Name: episodes episodes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.episodes
    ADD CONSTRAINT episodes_pkey PRIMARY KEY (id);


--
-- Name: episodes_topics episodes_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.episodes_topics
    ADD CONSTRAINT episodes_topics_pkey PRIMARY KEY (id);


--
-- Name: gbfm gbfm_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gbfm
    ADD CONSTRAINT gbfm_pkey PRIMARY KEY (id);


--
-- Name: gifm_authors gifm_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifm_authors
    ADD CONSTRAINT gifm_authors_pkey PRIMARY KEY (id);


--
-- Name: gifm_topics gifm_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gifm_topics
    ADD CONSTRAINT gifm_topics_pkey PRIMARY KEY (id);


--
-- Name: guides_authors guides_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.guides_authors
    ADD CONSTRAINT guides_authors_pkey PRIMARY KEY (id);


--
-- Name: guides guides_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.guides
    ADD CONSTRAINT guides_pkey PRIMARY KEY (id);


--
-- Name: guides_topics guides_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.guides_topics
    ADD CONSTRAINT guides_topics_pkey PRIMARY KEY (id);


--
-- Name: series_authors series_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_authors
    ADD CONSTRAINT series_authors_pkey PRIMARY KEY (id);


--
-- Name: series series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_pkey PRIMARY KEY (id);


--
-- Name: series_topics series_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series_topics
    ADD CONSTRAINT series_topics_pkey PRIMARY KEY (id);


--
-- Name: snacks_authors snacks_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.snacks_authors
    ADD CONSTRAINT snacks_authors_pkey PRIMARY KEY (id);


--
-- Name: snacks snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.snacks
    ADD CONSTRAINT snacks_pkey PRIMARY KEY (id);


--
-- Name: snacks_topics snacks_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.snacks_topics
    ADD CONSTRAINT snacks_topics_pkey PRIMARY KEY (id);


--
-- Name: topics topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--


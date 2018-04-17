--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Ubuntu 10.3-1)
-- Dumped by pg_dump version 10.3 (Ubuntu 10.3-1)

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
-- Name: authors_episodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors_episodes (
    id integer NOT NULL,
    episode_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors_episodes OWNER TO postgres;

--
-- Name: authors_episodes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_episodes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_episodes_id_seq OWNER TO postgres;

--
-- Name: authors_episodes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_episodes_id_seq OWNED BY public.authors_episodes.id;


--
-- Name: authors_gifm; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors_gifm (
    id integer NOT NULL,
    topic_id uuid NOT NULL,
    gifm_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors_gifm OWNER TO postgres;

--
-- Name: authors_gifm_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_gifm_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_gifm_id_seq OWNER TO postgres;

--
-- Name: authors_gifm_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_gifm_id_seq OWNED BY public.authors_gifm.id;


--
-- Name: authors_guides; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors_guides (
    id integer NOT NULL,
    guide_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors_guides OWNER TO postgres;

--
-- Name: authors_guides_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_guides_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_guides_id_seq OWNER TO postgres;

--
-- Name: authors_guides_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_guides_id_seq OWNED BY public.authors_guides.id;


--
-- Name: authors_series; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors_series (
    id integer NOT NULL,
    series_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors_series OWNER TO postgres;

--
-- Name: authors_series_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_series_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_series_id_seq OWNER TO postgres;

--
-- Name: authors_series_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_series_id_seq OWNED BY public.authors_series.id;


--
-- Name: authors_snacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors_snacks (
    id integer NOT NULL,
    snack_id uuid NOT NULL,
    author_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.authors_snacks OWNER TO postgres;

--
-- Name: authors_snacks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_snacks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_snacks_id_seq OWNER TO postgres;

--
-- Name: authors_snacks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_snacks_id_seq OWNED BY public.authors_snacks.id;


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
-- Name: topics_episodes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics_episodes (
    id integer NOT NULL,
    episode_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.topics_episodes OWNER TO postgres;

--
-- Name: topics_episodes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.topics_episodes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.topics_episodes_id_seq OWNER TO postgres;

--
-- Name: topics_episodes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.topics_episodes_id_seq OWNED BY public.topics_episodes.id;


--
-- Name: topics_gifm; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics_gifm (
    id integer NOT NULL,
    topic_id uuid NOT NULL,
    gifm_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.topics_gifm OWNER TO postgres;

--
-- Name: topics_gifm_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.topics_gifm_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.topics_gifm_id_seq OWNER TO postgres;

--
-- Name: topics_gifm_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.topics_gifm_id_seq OWNED BY public.topics_gifm.id;


--
-- Name: topics_guides; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics_guides (
    id integer NOT NULL,
    guide_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.topics_guides OWNER TO postgres;

--
-- Name: topics_guides_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.topics_guides_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.topics_guides_id_seq OWNER TO postgres;

--
-- Name: topics_guides_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.topics_guides_id_seq OWNED BY public.topics_guides.id;


--
-- Name: topics_series; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics_series (
    id integer NOT NULL,
    series_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.topics_series OWNER TO postgres;

--
-- Name: topics_series_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.topics_series_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.topics_series_id_seq OWNER TO postgres;

--
-- Name: topics_series_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.topics_series_id_seq OWNED BY public.topics_series.id;


--
-- Name: topics_snacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics_snacks (
    id integer NOT NULL,
    snack_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.topics_snacks OWNER TO postgres;

--
-- Name: topics_snacks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.topics_snacks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.topics_snacks_id_seq OWNER TO postgres;

--
-- Name: topics_snacks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.topics_snacks_id_seq OWNED BY public.topics_snacks.id;


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
-- Name: authors_episodes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_episodes ALTER COLUMN id SET DEFAULT nextval('public.authors_episodes_id_seq'::regclass);


--
-- Name: authors_gifm id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_gifm ALTER COLUMN id SET DEFAULT nextval('public.authors_gifm_id_seq'::regclass);


--
-- Name: authors_guides id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_guides ALTER COLUMN id SET DEFAULT nextval('public.authors_guides_id_seq'::regclass);


--
-- Name: authors_series id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_series ALTER COLUMN id SET DEFAULT nextval('public.authors_series_id_seq'::regclass);


--
-- Name: authors_snacks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_snacks ALTER COLUMN id SET DEFAULT nextval('public.authors_snacks_id_seq'::regclass);


--
-- Name: topics_episodes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_episodes ALTER COLUMN id SET DEFAULT nextval('public.topics_episodes_id_seq'::regclass);


--
-- Name: topics_gifm id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_gifm ALTER COLUMN id SET DEFAULT nextval('public.topics_gifm_id_seq'::regclass);


--
-- Name: topics_guides id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_guides ALTER COLUMN id SET DEFAULT nextval('public.topics_guides_id_seq'::regclass);


--
-- Name: topics_series id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_series ALTER COLUMN id SET DEFAULT nextval('public.topics_series_id_seq'::regclass);


--
-- Name: topics_snacks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_snacks ALTER COLUMN id SET DEFAULT nextval('public.topics_snacks_id_seq'::regclass);


--
-- Name: authors_episodes authors_episodes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_episodes
    ADD CONSTRAINT authors_episodes_pkey PRIMARY KEY (id);


--
-- Name: authors_gifm authors_gifm_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_gifm
    ADD CONSTRAINT authors_gifm_pkey PRIMARY KEY (id);


--
-- Name: authors_guides authors_guides_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_guides
    ADD CONSTRAINT authors_guides_pkey PRIMARY KEY (id);


--
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- Name: authors_series authors_series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_series
    ADD CONSTRAINT authors_series_pkey PRIMARY KEY (id);


--
-- Name: authors_snacks authors_snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors_snacks
    ADD CONSTRAINT authors_snacks_pkey PRIMARY KEY (id);


--
-- Name: episodes episodes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.episodes
    ADD CONSTRAINT episodes_pkey PRIMARY KEY (id);


--
-- Name: gbfm gbfm_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gbfm
    ADD CONSTRAINT gbfm_pkey PRIMARY KEY (id);


--
-- Name: guides guides_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.guides
    ADD CONSTRAINT guides_pkey PRIMARY KEY (id);


--
-- Name: series series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.series
    ADD CONSTRAINT series_pkey PRIMARY KEY (id);


--
-- Name: snacks snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.snacks
    ADD CONSTRAINT snacks_pkey PRIMARY KEY (id);


--
-- Name: topics_episodes topics_episodes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_episodes
    ADD CONSTRAINT topics_episodes_pkey PRIMARY KEY (id);


--
-- Name: topics_gifm topics_gifm_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_gifm
    ADD CONSTRAINT topics_gifm_pkey PRIMARY KEY (id);


--
-- Name: topics_guides topics_guides_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_guides
    ADD CONSTRAINT topics_guides_pkey PRIMARY KEY (id);


--
-- Name: topics topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (id);


--
-- Name: topics_series topics_series_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_series
    ADD CONSTRAINT topics_series_pkey PRIMARY KEY (id);


--
-- Name: topics_snacks topics_snacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics_snacks
    ADD CONSTRAINT topics_snacks_pkey PRIMARY KEY (id);


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


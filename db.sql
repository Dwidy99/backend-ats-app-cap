PGDMP                           z            compro    14.4    14.4 <    H           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            I           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            J           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            K           1262    41210    compro    DATABASE     j   CREATE DATABASE compro WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE compro;
                postgres    false            �            1259    41236    news    TABLE     e  CREATE TABLE public.news (
    id_news integer NOT NULL,
    id_user integer NOT NULL,
    id_category_news integer NOT NULL,
    updater character varying(255) NOT NULL,
    slug_news character varying(255) NOT NULL,
    judul_news character varying(255) NOT NULL,
    isi text NOT NULL,
    image character varying(255) NOT NULL,
    hits integer NOT NULL,
    status_news integer NOT NULL,
    jenis_news character varying(20) NOT NULL,
    keywords character varying(500) NOT NULL,
    date_post timestamp without time zone NOT NULL,
    date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.news;
       public         heap    postgres    false            �            1259    41235    berita_id_berita_seq    SEQUENCE     �   CREATE SEQUENCE public.berita_id_berita_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.berita_id_berita_seq;
       public          postgres    false    210            L           0    0    berita_id_berita_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.berita_id_berita_seq OWNED BY public.news.id_news;
          public          postgres    false    209            �            1259    41267    configuration    TABLE     �  CREATE TABLE public.configuration (
    id_configuration integer NOT NULL,
    id_user integer NOT NULL,
    webname character varying(50) NOT NULL,
    tagline character varying(100) DEFAULT NULL::character varying,
    email character varying(255) DEFAULT NULL::character varying,
    website character varying(255) DEFAULT NULL::character varying,
    phone character varying(25) DEFAULT NULL::character varying,
    address text,
    description character varying(300) DEFAULT NULL::character varying,
    keywords character varying(300) DEFAULT NULL::character varying,
    metatext text,
    map text,
    logo character varying(255) DEFAULT NULL::character varying,
    icon character varying(255) DEFAULT NULL::character varying,
    facebook character varying(255) DEFAULT NULL::character varying,
    instagram character varying(255) DEFAULT NULL::character varying,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
 !   DROP TABLE public.configuration;
       public         heap    postgres    false            �            1259    41288    contacts    TABLE     +  CREATE TABLE public.contacts (
    id_cantact integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    subject character varying(255) NOT NULL,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.contacts;
       public         heap    postgres    false            �            1259    41246    galeries    TABLE     �  CREATE TABLE public.galeries (
    id_galery integer NOT NULL,
    id_user integer NOT NULL,
    judul_galery character varying(255) NOT NULL,
    isi_galery text NOT NULL,
    website character varying(255) DEFAULT NULL::character varying,
    hits integer NOT NULL,
    image character varying(255) NOT NULL,
    galery_position integer NOT NULL,
    date_post timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.galeries;
       public         heap    postgres    false            �            1259    41245    galeri_id_galeri_seq    SEQUENCE     �   CREATE SEQUENCE public.galeri_id_galeri_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.galeri_id_galeri_seq;
       public          postgres    false    212            M           0    0    galeri_id_galeri_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.galeri_id_galeri_seq OWNED BY public.galeries.id_galery;
          public          postgres    false    211            �            1259    41257    news_category    TABLE        CREATE TABLE public.news_category (
    id_category_news integer NOT NULL,
    name_category character varying(255) NOT NULL,
    slug_category character varying(255) NOT NULL,
    "order" integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
 !   DROP TABLE public.news_category;
       public         heap    postgres    false            �            1259    41256 &   kategori_berita_id_kategori_berita_seq    SEQUENCE     �   CREATE SEQUENCE public.kategori_berita_id_kategori_berita_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 =   DROP SEQUENCE public.kategori_berita_id_kategori_berita_seq;
       public          postgres    false    214            N           0    0 &   kategori_berita_id_kategori_berita_seq    SEQUENCE OWNED BY     m   ALTER SEQUENCE public.kategori_berita_id_kategori_berita_seq OWNED BY public.news_category.id_category_news;
          public          postgres    false    213            �            1259    41266    konfigurasi_id_konfigurasi_seq    SEQUENCE     �   CREATE SEQUENCE public.konfigurasi_id_konfigurasi_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 5   DROP SEQUENCE public.konfigurasi_id_konfigurasi_seq;
       public          postgres    false    216            O           0    0    konfigurasi_id_konfigurasi_seq    SEQUENCE OWNED BY     e   ALTER SEQUENCE public.konfigurasi_id_konfigurasi_seq OWNED BY public.configuration.id_configuration;
          public          postgres    false    215            �            1259    41287    kontak_id_kontak_seq    SEQUENCE     �   CREATE SEQUENCE public.kontak_id_kontak_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.kontak_id_kontak_seq;
       public          postgres    false    218            P           0    0    kontak_id_kontak_seq    SEQUENCE OWNED BY     P   ALTER SEQUENCE public.kontak_id_kontak_seq OWNED BY public.contacts.id_cantact;
          public          postgres    false    217            �            1259    41298    services    TABLE     H  CREATE TABLE public.services (
    id_layanan integer NOT NULL,
    id_user integer NOT NULL,
    hits integer NOT NULL,
    title character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    content text NOT NULL,
    price character varying(255) NOT NULL,
    image character varying(255) NOT NULL,
    status character varying(25) NOT NULL,
    keywords character varying(500) DEFAULT NULL::character varying,
    date_post timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.services;
       public         heap    postgres    false            �            1259    41297    layanan_id_layanan_seq    SEQUENCE     �   CREATE SEQUENCE public.layanan_id_layanan_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.layanan_id_layanan_seq;
       public          postgres    false    220            Q           0    0    layanan_id_layanan_seq    SEQUENCE OWNED BY     R   ALTER SEQUENCE public.layanan_id_layanan_seq OWNED BY public.services.id_layanan;
          public          postgres    false    219            �            1259    41310    users    TABLE     h  CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    email character varying(200) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    level_access character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    41309    users_id_user_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_id_user_seq;
       public          postgres    false    222            R           0    0    users_id_user_seq    SEQUENCE OWNED BY     B   ALTER SEQUENCE public.users_id_user_seq OWNED BY public.users.id;
          public          postgres    false    221            �            1259    41320    users_token    TABLE     �   CREATE TABLE public.users_token (
    id_token integer NOT NULL,
    email character varying(255) NOT NULL,
    user_token character varying(255) NOT NULL,
    date_created timestamp with time zone NOT NULL
);
    DROP TABLE public.users_token;
       public         heap    postgres    false            �            1259    41319    users_token_id_token_seq    SEQUENCE     �   CREATE SEQUENCE public.users_token_id_token_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.users_token_id_token_seq;
       public          postgres    false    224            S           0    0    users_token_id_token_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.users_token_id_token_seq OWNED BY public.users_token.id_token;
          public          postgres    false    223            �           2604    41270    configuration id_configuration    DEFAULT     �   ALTER TABLE ONLY public.configuration ALTER COLUMN id_configuration SET DEFAULT nextval('public.konfigurasi_id_konfigurasi_seq'::regclass);
 M   ALTER TABLE public.configuration ALTER COLUMN id_configuration DROP DEFAULT;
       public          postgres    false    215    216    216            �           2604    41291    contacts id_cantact    DEFAULT     w   ALTER TABLE ONLY public.contacts ALTER COLUMN id_cantact SET DEFAULT nextval('public.kontak_id_kontak_seq'::regclass);
 B   ALTER TABLE public.contacts ALTER COLUMN id_cantact DROP DEFAULT;
       public          postgres    false    217    218    218            �           2604    41249    galeries id_galery    DEFAULT     v   ALTER TABLE ONLY public.galeries ALTER COLUMN id_galery SET DEFAULT nextval('public.galeri_id_galeri_seq'::regclass);
 A   ALTER TABLE public.galeries ALTER COLUMN id_galery DROP DEFAULT;
       public          postgres    false    211    212    212                       2604    41239    news id_news    DEFAULT     p   ALTER TABLE ONLY public.news ALTER COLUMN id_news SET DEFAULT nextval('public.berita_id_berita_seq'::regclass);
 ;   ALTER TABLE public.news ALTER COLUMN id_news DROP DEFAULT;
       public          postgres    false    209    210    210            �           2604    41260    news_category id_category_news    DEFAULT     �   ALTER TABLE ONLY public.news_category ALTER COLUMN id_category_news SET DEFAULT nextval('public.kategori_berita_id_kategori_berita_seq'::regclass);
 M   ALTER TABLE public.news_category ALTER COLUMN id_category_news DROP DEFAULT;
       public          postgres    false    214    213    214            �           2604    41301    services id_layanan    DEFAULT     y   ALTER TABLE ONLY public.services ALTER COLUMN id_layanan SET DEFAULT nextval('public.layanan_id_layanan_seq'::regclass);
 B   ALTER TABLE public.services ALTER COLUMN id_layanan DROP DEFAULT;
       public          postgres    false    219    220    220            �           2604    41313    users id    DEFAULT     i   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_user_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    222    221    222            �           2604    41323    users_token id_token    DEFAULT     |   ALTER TABLE ONLY public.users_token ALTER COLUMN id_token SET DEFAULT nextval('public.users_token_id_token_seq'::regclass);
 C   ALTER TABLE public.users_token ALTER COLUMN id_token DROP DEFAULT;
       public          postgres    false    223    224    224            =          0    41267    configuration 
   TABLE DATA           �   COPY public.configuration (id_configuration, id_user, webname, tagline, email, website, phone, address, description, keywords, metatext, map, logo, icon, facebook, instagram, created_at) FROM stdin;
    public          postgres    false    216   �M       ?          0    41288    contacts 
   TABLE DATA           Y   COPY public.contacts (id_cantact, name, email, subject, message, created_at) FROM stdin;
    public          postgres    false    218   N       9          0    41246    galeries 
   TABLE DATA           �   COPY public.galeries (id_galery, id_user, judul_galery, isi_galery, website, hits, image, galery_position, date_post, created_at) FROM stdin;
    public          postgres    false    212   /N       7          0    41236    news 
   TABLE DATA           �   COPY public.news (id_news, id_user, id_category_news, updater, slug_news, judul_news, isi, image, hits, status_news, jenis_news, keywords, date_post, date) FROM stdin;
    public          postgres    false    210   LN       ;          0    41257    news_category 
   TABLE DATA           l   COPY public.news_category (id_category_news, name_category, slug_category, "order", created_at) FROM stdin;
    public          postgres    false    214   iN       A          0    41298    services 
   TABLE DATA           �   COPY public.services (id_layanan, id_user, hits, title, slug, content, price, image, status, keywords, date_post, created_at) FROM stdin;
    public          postgres    false    220   �N       C          0    41310    users 
   TABLE DATA           ^   COPY public.users (id, name, email, username, password, level_access, created_at) FROM stdin;
    public          postgres    false    222   �N       E          0    41320    users_token 
   TABLE DATA           P   COPY public.users_token (id_token, email, user_token, date_created) FROM stdin;
    public          postgres    false    224   0O       T           0    0    berita_id_berita_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.berita_id_berita_seq', 1, false);
          public          postgres    false    209            U           0    0    galeri_id_galeri_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.galeri_id_galeri_seq', 1, false);
          public          postgres    false    211            V           0    0 &   kategori_berita_id_kategori_berita_seq    SEQUENCE SET     U   SELECT pg_catalog.setval('public.kategori_berita_id_kategori_berita_seq', 1, false);
          public          postgres    false    213            W           0    0    konfigurasi_id_konfigurasi_seq    SEQUENCE SET     M   SELECT pg_catalog.setval('public.konfigurasi_id_konfigurasi_seq', 1, false);
          public          postgres    false    215            X           0    0    kontak_id_kontak_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.kontak_id_kontak_seq', 1, false);
          public          postgres    false    217            Y           0    0    layanan_id_layanan_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.layanan_id_layanan_seq', 1, false);
          public          postgres    false    219            Z           0    0    users_id_user_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.users_id_user_seq', 4, true);
          public          postgres    false    221            [           0    0    users_token_id_token_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.users_token_id_token_seq', 1, false);
          public          postgres    false    223            �           2606    41244    news berita_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.news
    ADD CONSTRAINT berita_pkey PRIMARY KEY (id_news);
 :   ALTER TABLE ONLY public.news DROP CONSTRAINT berita_pkey;
       public            postgres    false    210            �           2606    41255    galeries galeri_pkey 
   CONSTRAINT     Y   ALTER TABLE ONLY public.galeries
    ADD CONSTRAINT galeri_pkey PRIMARY KEY (id_galery);
 >   ALTER TABLE ONLY public.galeries DROP CONSTRAINT galeri_pkey;
       public            postgres    false    212            �           2606    41265 "   news_category kategori_berita_pkey 
   CONSTRAINT     n   ALTER TABLE ONLY public.news_category
    ADD CONSTRAINT kategori_berita_pkey PRIMARY KEY (id_category_news);
 L   ALTER TABLE ONLY public.news_category DROP CONSTRAINT kategori_berita_pkey;
       public            postgres    false    214            �           2606    41285    configuration konfigurasi_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.configuration
    ADD CONSTRAINT konfigurasi_pkey PRIMARY KEY (id_configuration);
 H   ALTER TABLE ONLY public.configuration DROP CONSTRAINT konfigurasi_pkey;
       public            postgres    false    216            �           2606    41296    contacts kontak_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.contacts
    ADD CONSTRAINT kontak_pkey PRIMARY KEY (id_cantact);
 >   ALTER TABLE ONLY public.contacts DROP CONSTRAINT kontak_pkey;
       public            postgres    false    218            �           2606    41308    services layanan_pkey 
   CONSTRAINT     [   ALTER TABLE ONLY public.services
    ADD CONSTRAINT layanan_pkey PRIMARY KEY (id_layanan);
 ?   ALTER TABLE ONLY public.services DROP CONSTRAINT layanan_pkey;
       public            postgres    false    220            �           2606    41318    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    222            �           2606    41327    users_token users_token_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.users_token
    ADD CONSTRAINT users_token_pkey PRIMARY KEY (id_token);
 F   ALTER TABLE ONLY public.users_token DROP CONSTRAINT users_token_pkey;
       public            postgres    false    224            =      x������ � �      ?      x������ � �      9      x������ � �      7      x������ � �      ;      x������ � �      A      x������ � �      C   }   x�3��L)�tH�M���K���8U�ULT��K|�rL=\KKK�-r����J�}}�|�S�rBM,S��½�M}�S-�8Sr3�8���t,u����L�L,��,��-�b���� �g!      E      x������ � �     
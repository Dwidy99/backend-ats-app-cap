PGDMP     #        
            z            compro    14.4    14.4 <    H           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            I           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            J           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            K           1262    41210    compro    DATABASE     j   CREATE DATABASE compro WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE compro;
                postgres    false            �            1259    41236    berita    TABLE     �  CREATE TABLE public.berita (
    id_berita integer NOT NULL,
    id_user integer NOT NULL,
    id_kategori_berita integer NOT NULL,
    updater character varying(255) NOT NULL,
    slug_berita character varying(255) NOT NULL,
    judul_berita character varying(255) NOT NULL,
    isi text NOT NULL,
    gambar character varying(255) NOT NULL,
    hits integer NOT NULL,
    status_berita integer NOT NULL,
    jenis_berita character varying(20) NOT NULL,
    keywords character varying(500) NOT NULL,
    tanggal_mulai date,
    tanggal_selesai date,
    tanggal_post timestamp without time zone NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.berita;
       public         heap    postgres    false            �            1259    41235    berita_id_berita_seq    SEQUENCE     �   CREATE SEQUENCE public.berita_id_berita_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.berita_id_berita_seq;
       public          postgres    false    210            L           0    0    berita_id_berita_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.berita_id_berita_seq OWNED BY public.berita.id_berita;
          public          postgres    false    209            �            1259    41246    galeri    TABLE     �  CREATE TABLE public.galeri (
    id_galeri integer NOT NULL,
    id_user integer NOT NULL,
    judul_galeri character varying(255) NOT NULL,
    isi_galeri text NOT NULL,
    website character varying(255) DEFAULT NULL::character varying,
    hits integer NOT NULL,
    gambar character varying(255) NOT NULL,
    posisi_galeri integer NOT NULL,
    tanggal_post timestamp without time zone NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.galeri;
       public         heap    postgres    false            �            1259    41245    galeri_id_galeri_seq    SEQUENCE     �   CREATE SEQUENCE public.galeri_id_galeri_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.galeri_id_galeri_seq;
       public          postgres    false    212            M           0    0    galeri_id_galeri_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.galeri_id_galeri_seq OWNED BY public.galeri.id_galeri;
          public          postgres    false    211            �            1259    41257    kategori_berita    TABLE        CREATE TABLE public.kategori_berita (
    id_kategori_berita integer NOT NULL,
    nama_kategori character varying(255) NOT NULL,
    slug_kategori character varying(255) NOT NULL,
    urutan integer NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
 #   DROP TABLE public.kategori_berita;
       public         heap    postgres    false            �            1259    41256 &   kategori_berita_id_kategori_berita_seq    SEQUENCE     �   CREATE SEQUENCE public.kategori_berita_id_kategori_berita_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 =   DROP SEQUENCE public.kategori_berita_id_kategori_berita_seq;
       public          postgres    false    214            N           0    0 &   kategori_berita_id_kategori_berita_seq    SEQUENCE OWNED BY     q   ALTER SEQUENCE public.kategori_berita_id_kategori_berita_seq OWNED BY public.kategori_berita.id_kategori_berita;
          public          postgres    false    213            �            1259    41267    konfigurasi    TABLE     �  CREATE TABLE public.konfigurasi (
    id_konfigurasi integer NOT NULL,
    id_user integer NOT NULL,
    namaweb character varying(50) NOT NULL,
    tagline character varying(100) DEFAULT NULL::character varying,
    email character varying(255) DEFAULT NULL::character varying,
    website character varying(255) DEFAULT NULL::character varying,
    telepon character varying(25) DEFAULT NULL::character varying,
    alamat text,
    deskripsi character varying(300) DEFAULT NULL::character varying,
    keywords character varying(300) DEFAULT NULL::character varying,
    metatext text,
    map text,
    logo character varying(255) DEFAULT NULL::character varying,
    icon character varying(255) DEFAULT NULL::character varying,
    facebook character varying(255) DEFAULT NULL::character varying,
    instagram character varying(255) DEFAULT NULL::character varying,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.konfigurasi;
       public         heap    postgres    false            �            1259    41266    konfigurasi_id_konfigurasi_seq    SEQUENCE     �   CREATE SEQUENCE public.konfigurasi_id_konfigurasi_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 5   DROP SEQUENCE public.konfigurasi_id_konfigurasi_seq;
       public          postgres    false    216            O           0    0    konfigurasi_id_konfigurasi_seq    SEQUENCE OWNED BY     a   ALTER SEQUENCE public.konfigurasi_id_konfigurasi_seq OWNED BY public.konfigurasi.id_konfigurasi;
          public          postgres    false    215            �            1259    41288    kontak    TABLE     &  CREATE TABLE public.kontak (
    id_kontak integer NOT NULL,
    nama character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    subject character varying(255) NOT NULL,
    pesan text NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.kontak;
       public         heap    postgres    false            �            1259    41287    kontak_id_kontak_seq    SEQUENCE     �   CREATE SEQUENCE public.kontak_id_kontak_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.kontak_id_kontak_seq;
       public          postgres    false    218            P           0    0    kontak_id_kontak_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.kontak_id_kontak_seq OWNED BY public.kontak.id_kontak;
          public          postgres    false    217            �            1259    41298    layanan    TABLE     j  CREATE TABLE public.layanan (
    id_layanan integer NOT NULL,
    id_user integer NOT NULL,
    hits integer NOT NULL,
    judul_layanan character varying(255) NOT NULL,
    slug_layanan character varying(255) NOT NULL,
    isi_layanan text NOT NULL,
    harga character varying(255) NOT NULL,
    gambar character varying(255) NOT NULL,
    status_layanan character varying(25) NOT NULL,
    keywords character varying(500) DEFAULT NULL::character varying,
    tanggal_post timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);
    DROP TABLE public.layanan;
       public         heap    postgres    false            �            1259    41297    layanan_id_layanan_seq    SEQUENCE     �   CREATE SEQUENCE public.layanan_id_layanan_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.layanan_id_layanan_seq;
       public          postgres    false    220            Q           0    0    layanan_id_layanan_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.layanan_id_layanan_seq OWNED BY public.layanan.id_layanan;
          public          postgres    false    219            �            1259    41310    users    TABLE     i  CREATE TABLE public.users (
    id_user integer NOT NULL,
    nama character varying(50) NOT NULL,
    email character varying(200) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    akses_level character varying(255) NOT NULL,
    tanggal timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
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
       public          postgres    false    222            R           0    0    users_id_user_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.users_id_user_seq OWNED BY public.users.id_user;
          public          postgres    false    221            �            1259    41320    users_token    TABLE     �   CREATE TABLE public.users_token (
    id_token integer NOT NULL,
    email character varying(255) NOT NULL,
    user_token character varying(255) NOT NULL,
    tanggal_buat integer NOT NULL
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
          public          postgres    false    223                       2604    41239    berita id_berita    DEFAULT     t   ALTER TABLE ONLY public.berita ALTER COLUMN id_berita SET DEFAULT nextval('public.berita_id_berita_seq'::regclass);
 ?   ALTER TABLE public.berita ALTER COLUMN id_berita DROP DEFAULT;
       public          postgres    false    210    209    210            �           2604    41249    galeri id_galeri    DEFAULT     t   ALTER TABLE ONLY public.galeri ALTER COLUMN id_galeri SET DEFAULT nextval('public.galeri_id_galeri_seq'::regclass);
 ?   ALTER TABLE public.galeri ALTER COLUMN id_galeri DROP DEFAULT;
       public          postgres    false    211    212    212            �           2604    41260 "   kategori_berita id_kategori_berita    DEFAULT     �   ALTER TABLE ONLY public.kategori_berita ALTER COLUMN id_kategori_berita SET DEFAULT nextval('public.kategori_berita_id_kategori_berita_seq'::regclass);
 Q   ALTER TABLE public.kategori_berita ALTER COLUMN id_kategori_berita DROP DEFAULT;
       public          postgres    false    214    213    214            �           2604    41270    konfigurasi id_konfigurasi    DEFAULT     �   ALTER TABLE ONLY public.konfigurasi ALTER COLUMN id_konfigurasi SET DEFAULT nextval('public.konfigurasi_id_konfigurasi_seq'::regclass);
 I   ALTER TABLE public.konfigurasi ALTER COLUMN id_konfigurasi DROP DEFAULT;
       public          postgres    false    216    215    216            �           2604    41291    kontak id_kontak    DEFAULT     t   ALTER TABLE ONLY public.kontak ALTER COLUMN id_kontak SET DEFAULT nextval('public.kontak_id_kontak_seq'::regclass);
 ?   ALTER TABLE public.kontak ALTER COLUMN id_kontak DROP DEFAULT;
       public          postgres    false    218    217    218            �           2604    41301    layanan id_layanan    DEFAULT     x   ALTER TABLE ONLY public.layanan ALTER COLUMN id_layanan SET DEFAULT nextval('public.layanan_id_layanan_seq'::regclass);
 A   ALTER TABLE public.layanan ALTER COLUMN id_layanan DROP DEFAULT;
       public          postgres    false    219    220    220            �           2604    41313    users id_user    DEFAULT     n   ALTER TABLE ONLY public.users ALTER COLUMN id_user SET DEFAULT nextval('public.users_id_user_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN id_user DROP DEFAULT;
       public          postgres    false    222    221    222            �           2604    41323    users_token id_token    DEFAULT     |   ALTER TABLE ONLY public.users_token ALTER COLUMN id_token SET DEFAULT nextval('public.users_token_id_token_seq'::regclass);
 C   ALTER TABLE public.users_token ALTER COLUMN id_token DROP DEFAULT;
       public          postgres    false    224    223    224            7          0    41236    berita 
   TABLE DATA           �   COPY public.berita (id_berita, id_user, id_kategori_berita, updater, slug_berita, judul_berita, isi, gambar, hits, status_berita, jenis_berita, keywords, tanggal_mulai, tanggal_selesai, tanggal_post, tanggal) FROM stdin;
    public          postgres    false    210   �N       9          0    41246    galeri 
   TABLE DATA           �   COPY public.galeri (id_galeri, id_user, judul_galeri, isi_galeri, website, hits, gambar, posisi_galeri, tanggal_post, tanggal) FROM stdin;
    public          postgres    false    212   �N       ;          0    41257    kategori_berita 
   TABLE DATA           l   COPY public.kategori_berita (id_kategori_berita, nama_kategori, slug_kategori, urutan, tanggal) FROM stdin;
    public          postgres    false    214   �N       =          0    41267    konfigurasi 
   TABLE DATA           �   COPY public.konfigurasi (id_konfigurasi, id_user, namaweb, tagline, email, website, telepon, alamat, deskripsi, keywords, metatext, map, logo, icon, facebook, instagram, tanggal) FROM stdin;
    public          postgres    false    216   �N       ?          0    41288    kontak 
   TABLE DATA           Q   COPY public.kontak (id_kontak, nama, email, subject, pesan, tanggal) FROM stdin;
    public          postgres    false    218   �N       A          0    41298    layanan 
   TABLE DATA           �   COPY public.layanan (id_layanan, id_user, hits, judul_layanan, slug_layanan, isi_layanan, harga, gambar, status_layanan, keywords, tanggal_post, tanggal) FROM stdin;
    public          postgres    false    220   O       C          0    41310    users 
   TABLE DATA           _   COPY public.users (id_user, nama, email, username, password, akses_level, tanggal) FROM stdin;
    public          postgres    false    222   5O       E          0    41320    users_token 
   TABLE DATA           P   COPY public.users_token (id_token, email, user_token, tanggal_buat) FROM stdin;
    public          postgres    false    224   RO       T           0    0    berita_id_berita_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.berita_id_berita_seq', 1, false);
          public          postgres    false    209            U           0    0    galeri_id_galeri_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.galeri_id_galeri_seq', 1, false);
          public          postgres    false    211            V           0    0 &   kategori_berita_id_kategori_berita_seq    SEQUENCE SET     U   SELECT pg_catalog.setval('public.kategori_berita_id_kategori_berita_seq', 1, false);
          public          postgres    false    213            W           0    0    konfigurasi_id_konfigurasi_seq    SEQUENCE SET     M   SELECT pg_catalog.setval('public.konfigurasi_id_konfigurasi_seq', 1, false);
          public          postgres    false    215            X           0    0    kontak_id_kontak_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.kontak_id_kontak_seq', 1, false);
          public          postgres    false    217            Y           0    0    layanan_id_layanan_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.layanan_id_layanan_seq', 1, false);
          public          postgres    false    219            Z           0    0    users_id_user_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.users_id_user_seq', 1, false);
          public          postgres    false    221            [           0    0    users_token_id_token_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.users_token_id_token_seq', 1, false);
          public          postgres    false    223            �           2606    41244    berita berita_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.berita
    ADD CONSTRAINT berita_pkey PRIMARY KEY (id_berita);
 <   ALTER TABLE ONLY public.berita DROP CONSTRAINT berita_pkey;
       public            postgres    false    210            �           2606    41255    galeri galeri_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.galeri
    ADD CONSTRAINT galeri_pkey PRIMARY KEY (id_galeri);
 <   ALTER TABLE ONLY public.galeri DROP CONSTRAINT galeri_pkey;
       public            postgres    false    212            �           2606    41265 $   kategori_berita kategori_berita_pkey 
   CONSTRAINT     r   ALTER TABLE ONLY public.kategori_berita
    ADD CONSTRAINT kategori_berita_pkey PRIMARY KEY (id_kategori_berita);
 N   ALTER TABLE ONLY public.kategori_berita DROP CONSTRAINT kategori_berita_pkey;
       public            postgres    false    214            �           2606    41285    konfigurasi konfigurasi_pkey 
   CONSTRAINT     f   ALTER TABLE ONLY public.konfigurasi
    ADD CONSTRAINT konfigurasi_pkey PRIMARY KEY (id_konfigurasi);
 F   ALTER TABLE ONLY public.konfigurasi DROP CONSTRAINT konfigurasi_pkey;
       public            postgres    false    216            �           2606    41296    kontak kontak_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.kontak
    ADD CONSTRAINT kontak_pkey PRIMARY KEY (id_kontak);
 <   ALTER TABLE ONLY public.kontak DROP CONSTRAINT kontak_pkey;
       public            postgres    false    218            �           2606    41308    layanan layanan_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.layanan
    ADD CONSTRAINT layanan_pkey PRIMARY KEY (id_layanan);
 >   ALTER TABLE ONLY public.layanan DROP CONSTRAINT layanan_pkey;
       public            postgres    false    220            �           2606    41318    users users_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id_user);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    222            �           2606    41327    users_token users_token_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.users_token
    ADD CONSTRAINT users_token_pkey PRIMARY KEY (id_token);
 F   ALTER TABLE ONLY public.users_token DROP CONSTRAINT users_token_pkey;
       public            postgres    false    224            7      x������ � �      9      x������ � �      ;      x������ � �      =      x������ � �      ?      x������ � �      A      x������ � �      C      x������ � �      E      x������ � �     
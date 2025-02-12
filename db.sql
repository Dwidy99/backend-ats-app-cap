PGDMP     1            	        z            ats_app    14.4    14.4 @    >           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            @           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            A           1262    41884    ats_app    DATABASE     k   CREATE DATABASE ats_app WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE ats_app;
                postgres    false            �            1259    41972 	   applicant    TABLE     W  CREATE TABLE public.applicant (
    id integer NOT NULL,
    user_id integer,
    first_name character varying(255),
    last_name character varying(255),
    avatar character varying(255),
    account_status integer,
    last_education character varying(255),
    linkedin_url character varying(255),
    github_url character varying(255)
);
    DROP TABLE public.applicant;
       public         heap    postgres    false            �            1259    41971    Applicant_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Applicant_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public."Applicant_id_seq";
       public          postgres    false    210            B           0    0    Applicant_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public."Applicant_id_seq" OWNED BY public.applicant.id;
          public          postgres    false    209            �            1259    42004 	   companies    TABLE       CREATE TABLE public.companies (
    id integer NOT NULL,
    name character varying(255),
    email character varying(255),
    address text,
    contact character varying(255),
    website character varying(255),
    created_at timestamp without time zone
);
    DROP TABLE public.companies;
       public         heap    postgres    false            �            1259    42003    Company_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Company_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public."Company_id_seq";
       public          postgres    false    216            C           0    0    Company_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public."Company_id_seq" OWNED BY public.companies.id;
          public          postgres    false    215            �            1259    42026 	   employees    TABLE     �   CREATE TABLE public.employees (
    id integer NOT NULL,
    user_id integer,
    company_id integer,
    name character varying(255),
    contact character varying(255)
);
    DROP TABLE public.employees;
       public         heap    postgres    false            �            1259    42025    Employee_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Employee_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."Employee_id_seq";
       public          postgres    false    222            D           0    0    Employee_id_seq    SEQUENCE OWNED BY     F   ALTER SEQUENCE public."Employee_id_seq" OWNED BY public.employees.id;
          public          postgres    false    221            �            1259    41988    jobapplications    TABLE     �   CREATE TABLE public.jobapplications (
    id integer NOT NULL,
    applicant_id integer,
    job_id integer,
    status character varying(255),
    created_at timestamp without time zone
);
 #   DROP TABLE public.jobapplications;
       public         heap    postgres    false            �            1259    41987    JobApplication_id_seq    SEQUENCE     �   CREATE SEQUENCE public."JobApplication_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public."JobApplication_id_seq";
       public          postgres    false    212            E           0    0    JobApplication_id_seq    SEQUENCE OWNED BY     R   ALTER SEQUENCE public."JobApplication_id_seq" OWNED BY public.jobapplications.id;
          public          postgres    false    211            �            1259    42035    jobexperience    TABLE     �   CREATE TABLE public.jobexperience (
    id integer NOT NULL,
    applicant_id integer,
    company_name character varying(255),
    role character varying(255),
    description text,
    "dateStart" date,
    "dateEnd" date,
    status integer
);
 !   DROP TABLE public.jobexperience;
       public         heap    postgres    false            �            1259    42034    JobExperience_id_seq    SEQUENCE     �   CREATE SEQUENCE public."JobExperience_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public."JobExperience_id_seq";
       public          postgres    false    224            F           0    0    JobExperience_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public."JobExperience_id_seq" OWNED BY public.jobexperience.id;
          public          postgres    false    223            �            1259    42013 	   jobskills    TABLE     \   CREATE TABLE public.jobskills (
    id integer NOT NULL,
    name character varying(255)
);
    DROP TABLE public.jobskills;
       public         heap    postgres    false            �            1259    42012    JobSkills_id_seq    SEQUENCE     �   CREATE SEQUENCE public."JobSkills_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public."JobSkills_id_seq";
       public          postgres    false    218            G           0    0    JobSkills_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public."JobSkills_id_seq" OWNED BY public.jobskills.id;
          public          postgres    false    217            �            1259    41995    jobs    TABLE     �  CREATE TABLE public.jobs (
    id integer NOT NULL,
    company_id integer,
    title character varying(255),
    description text,
    location character varying(255),
    salary double precision,
    type character varying(255),
    level_of_experience character varying(255),
    date_start date,
    date_end date,
    created_at timestamp without time zone,
    posted_by integer
);
    DROP TABLE public.jobs;
       public         heap    postgres    false            �            1259    41994    Jobs_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Jobs_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public."Jobs_id_seq";
       public          postgres    false    214            H           0    0    Jobs_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public."Jobs_id_seq" OWNED BY public.jobs.id;
          public          postgres    false    213            �            1259    42022    jobskillapplicants    TABLE     [   CREATE TABLE public.jobskillapplicants (
    id_skill integer,
    id_applicant integer
);
 &   DROP TABLE public.jobskillapplicants;
       public         heap    postgres    false            �            1259    42019    jobskillrequirements    TABLE     W   CREATE TABLE public.jobskillrequirements (
    id_skill integer,
    id_job integer
);
 (   DROP TABLE public.jobskillrequirements;
       public         heap    postgres    false            �            1259    42050    users    TABLE       CREATE TABLE public.users (
    id bigint NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character(255) NOT NULL,
    role character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    42049    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    226            I           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    225            �           2604    41975    applicant id    DEFAULT     n   ALTER TABLE ONLY public.applicant ALTER COLUMN id SET DEFAULT nextval('public."Applicant_id_seq"'::regclass);
 ;   ALTER TABLE public.applicant ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    210    209    210            �           2604    42007    companies id    DEFAULT     l   ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public."Company_id_seq"'::regclass);
 ;   ALTER TABLE public.companies ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    216    216            �           2604    42029    employees id    DEFAULT     m   ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public."Employee_id_seq"'::regclass);
 ;   ALTER TABLE public.employees ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    222    222            �           2604    41991    jobapplications id    DEFAULT     y   ALTER TABLE ONLY public.jobapplications ALTER COLUMN id SET DEFAULT nextval('public."JobApplication_id_seq"'::regclass);
 A   ALTER TABLE public.jobapplications ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    212    211    212            �           2604    42038    jobexperience id    DEFAULT     v   ALTER TABLE ONLY public.jobexperience ALTER COLUMN id SET DEFAULT nextval('public."JobExperience_id_seq"'::regclass);
 ?   ALTER TABLE public.jobexperience ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    224    223    224            �           2604    41998    jobs id    DEFAULT     d   ALTER TABLE ONLY public.jobs ALTER COLUMN id SET DEFAULT nextval('public."Jobs_id_seq"'::regclass);
 6   ALTER TABLE public.jobs ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    213    214    214            �           2604    42016    jobskills id    DEFAULT     n   ALTER TABLE ONLY public.jobskills ALTER COLUMN id SET DEFAULT nextval('public."JobSkills_id_seq"'::regclass);
 ;   ALTER TABLE public.jobskills ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    217    218            �           2604    42053    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    225    226    226            +          0    41972 	   applicant 
   TABLE DATA           �   COPY public.applicant (id, user_id, first_name, last_name, avatar, account_status, last_education, linkedin_url, github_url) FROM stdin;
    public          postgres    false    210   G       1          0    42004 	   companies 
   TABLE DATA           [   COPY public.companies (id, name, email, address, contact, website, created_at) FROM stdin;
    public          postgres    false    216   G       7          0    42026 	   employees 
   TABLE DATA           K   COPY public.employees (id, user_id, company_id, name, contact) FROM stdin;
    public          postgres    false    222   <G       -          0    41988    jobapplications 
   TABLE DATA           W   COPY public.jobapplications (id, applicant_id, job_id, status, created_at) FROM stdin;
    public          postgres    false    212   YG       9          0    42035    jobexperience 
   TABLE DATA           z   COPY public.jobexperience (id, applicant_id, company_name, role, description, "dateStart", "dateEnd", status) FROM stdin;
    public          postgres    false    224   vG       /          0    41995    jobs 
   TABLE DATA           �   COPY public.jobs (id, company_id, title, description, location, salary, type, level_of_experience, date_start, date_end, created_at, posted_by) FROM stdin;
    public          postgres    false    214   �G       5          0    42022    jobskillapplicants 
   TABLE DATA           D   COPY public.jobskillapplicants (id_skill, id_applicant) FROM stdin;
    public          postgres    false    220   �G       4          0    42019    jobskillrequirements 
   TABLE DATA           @   COPY public.jobskillrequirements (id_skill, id_job) FROM stdin;
    public          postgres    false    219   �G       3          0    42013 	   jobskills 
   TABLE DATA           -   COPY public.jobskills (id, name) FROM stdin;
    public          postgres    false    218   �G       ;          0    42050    users 
   TABLE DATA           P   COPY public.users (id, username, email, password, role, created_at) FROM stdin;
    public          postgres    false    226   H       J           0    0    Applicant_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public."Applicant_id_seq"', 1, false);
          public          postgres    false    209            K           0    0    Company_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public."Company_id_seq"', 1, false);
          public          postgres    false    215            L           0    0    Employee_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public."Employee_id_seq"', 1, false);
          public          postgres    false    221            M           0    0    JobApplication_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public."JobApplication_id_seq"', 1, false);
          public          postgres    false    211            N           0    0    JobExperience_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public."JobExperience_id_seq"', 1, false);
          public          postgres    false    223            O           0    0    JobSkills_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public."JobSkills_id_seq"', 1, false);
          public          postgres    false    217            P           0    0    Jobs_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public."Jobs_id_seq"', 1, false);
          public          postgres    false    213            Q           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 5, true);
          public          postgres    false    225            �           2606    41979    applicant Applicant_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.applicant
    ADD CONSTRAINT "Applicant_pkey" PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.applicant DROP CONSTRAINT "Applicant_pkey";
       public            postgres    false    210            �           2606    42011    companies Company_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.companies
    ADD CONSTRAINT "Company_pkey" PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.companies DROP CONSTRAINT "Company_pkey";
       public            postgres    false    216            �           2606    42033    employees Employee_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.employees
    ADD CONSTRAINT "Employee_pkey" PRIMARY KEY (id);
 C   ALTER TABLE ONLY public.employees DROP CONSTRAINT "Employee_pkey";
       public            postgres    false    222            �           2606    41993 #   jobapplications JobApplication_pkey 
   CONSTRAINT     c   ALTER TABLE ONLY public.jobapplications
    ADD CONSTRAINT "JobApplication_pkey" PRIMARY KEY (id);
 O   ALTER TABLE ONLY public.jobapplications DROP CONSTRAINT "JobApplication_pkey";
       public            postgres    false    212            �           2606    42042     jobexperience JobExperience_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.jobexperience
    ADD CONSTRAINT "JobExperience_pkey" PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.jobexperience DROP CONSTRAINT "JobExperience_pkey";
       public            postgres    false    224            �           2606    42018    jobskills JobSkills_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.jobskills
    ADD CONSTRAINT "JobSkills_pkey" PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.jobskills DROP CONSTRAINT "JobSkills_pkey";
       public            postgres    false    218            �           2606    42002    jobs Jobs_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT "Jobs_pkey" PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.jobs DROP CONSTRAINT "Jobs_pkey";
       public            postgres    false    214            �           2606    42057    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    226            +      x������ � �      1      x������ � �      7      x������ � �      -      x������ � �      9      x������ � �      /      x������ � �      5      x������ � �      4      x������ � �      3      x������ � �      ;     x��йn�@���]����rU���$`�(Di�S �9-��'�GI:KV~�FSL���:iY\���y��,ߡh_raH�:+��d.+�F�ľ��D�X���&��\1w�E��/�ΖO���s��A��:êX�"�n@�	w�����aY9t���~0C�h��Ce��ӏC��&u�A1C��_ �ov�C)"�$�$n��k��Km��:�U�2�?$�C���S|>�Ǫ<���ˊ�p�4�^6׮}��L��Љ�((X��p����l��     
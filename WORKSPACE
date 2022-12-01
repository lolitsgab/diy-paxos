load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "099a9fb96a376ccbbb7d291ed4ecbdfd42f6bc822ab77ae6f1b5cb9e914e94fa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "efbbba6ac1a4fd342d5122cbdfdb82aeb2cf2862e35022c752eaddffada7c3f3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:hzAQntlaYRkVSFEfj9OTWlVV1H155FMD8BTKktLv0QI=",
    version = "v0.0.0-20210930031921-04548b0d99d4",
)

go_repository(
    name = "com_github_cncf_xds_go",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:zH8ljVhhq7yC0MIeUL/IviMtY8hx2mK8cN9wEYb8ggw=",
    version = "v0.0.0-20211011173535-cb28da3451f1",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:xvqufLtNVwAhN8NMyWklVgxnWohi+wtMGQMhtxexlm0=",
    version = "v0.10.2-0.20220325020618-49ff273808a1",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:BKbKCqvP6I+rmFHt06ZmyQtvB8xAkWdhFyr0ZUNZcxQ=",
    version = "v0.5.6",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:EVhdT+1Kseyi1/pUmXKaFxYsDNy9RQYkMWRH68J/W7Y=",
    version = "v1.1.2",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:eOI3/cP2VTU6uZLDYAoic+eyzzB9YyGmJ7eIjl8rOPg=",
    version = "v0.34.0",
)

go_repository(
    name = "com_google_cloud_go_accessapproval",
    importpath = "cloud.google.com/go/accessapproval",
    sum = "h1:/nTivgnV/n1CaAeo+ekGexTYUsKEU9jUVkoY5359+3Q=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_accesscontextmanager",
    importpath = "cloud.google.com/go/accesscontextmanager",
    sum = "h1:CFhNhU7pcD11cuDkQdrE6PQJgv0EXNKNv06jIzbLlCU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_aiplatform",
    importpath = "cloud.google.com/go/aiplatform",
    sum = "h1:QqHZT1IMldf/daXoSnkJWBIqGBsw50X+xP6HSVzLRPo=",
    version = "v1.24.0",
)

go_repository(
    name = "com_google_cloud_go_analytics",
    importpath = "cloud.google.com/go/analytics",
    sum = "h1:NKw6PpQi6V1O+KsjuTd+bhip9d0REYu4NevC45vtGp8=",
    version = "v0.12.0",
)

go_repository(
    name = "com_google_cloud_go_apigateway",
    importpath = "cloud.google.com/go/apigateway",
    sum = "h1:IIoXKR7FKrEAQhMTz5hK2wiDz2WNFHS7eVr/L1lE/rM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_apigeeconnect",
    importpath = "cloud.google.com/go/apigeeconnect",
    sum = "h1:AONoTYJviyv1vS4IkvWzq69gEVdvHx35wKXc+e6wjZQ=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_appengine",
    importpath = "cloud.google.com/go/appengine",
    sum = "h1:lmG+O5oaR9xNwaRBwE2XoMhwQHsHql5IoiGr1ptdDwU=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_area120",
    importpath = "cloud.google.com/go/area120",
    sum = "h1:TCMhwWEWhCn8d44/Zs7UCICTWje9j3HuV6nVGMjdpYw=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_artifactregistry",
    importpath = "cloud.google.com/go/artifactregistry",
    sum = "h1:3d0LRAU1K6vfqCahhl9fx2oGHcq+s5gftdix4v8Ibrc=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_asset",
    importpath = "cloud.google.com/go/asset",
    sum = "h1:aCrlaLGJWTODJX4G56ZYzJefITKEWNfbjjtHSzWpxW0=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_assuredworkloads",
    importpath = "cloud.google.com/go/assuredworkloads",
    sum = "h1:hhIdCOowsT1GG5eMCIA0OwK6USRuYTou/1ZeNxCSRtA=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_automl",
    importpath = "cloud.google.com/go/automl",
    sum = "h1:BMioyXSbg7d7xLibn47cs0elW6RT780IUWr42W8rp2Q=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_baremetalsolution",
    importpath = "cloud.google.com/go/baremetalsolution",
    sum = "h1:g9KO6SkakcYPcc/XjAzeuUrEOXlYPnMpuiaywYaGrmQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_batch",
    importpath = "cloud.google.com/go/batch",
    sum = "h1:1jvEBY55OH4Sd2FxEXQfxGExFWov1A/IaRe+Z5Z71Fw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_beyondcorp",
    importpath = "cloud.google.com/go/beyondcorp",
    sum = "h1:w+4kThysgl0JiKshi2MKDCg2NZgOyqOI0wq2eBZyrzA=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:u0fvz5ysJBe1jwUPI4LuPwAX+o+6fCUwf3ECeg6eDUQ=",
    version = "v1.43.0",
)

go_repository(
    name = "com_google_cloud_go_billing",
    importpath = "cloud.google.com/go/billing",
    sum = "h1:Xkii76HWELHwBtkQVZvqmSo9GTr0O+tIbRNnMcGdlg4=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_binaryauthorization",
    importpath = "cloud.google.com/go/binaryauthorization",
    sum = "h1:pL70vXWn9TitQYXBWTK2abHl2JHLwkFRjYw6VflRqEA=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_certificatemanager",
    importpath = "cloud.google.com/go/certificatemanager",
    sum = "h1:tzbR4UHBbgsewMWUD93JHi8EBi/gHBoSAcY1/sThFGk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_channel",
    importpath = "cloud.google.com/go/channel",
    sum = "h1:pNuUlZx0Jb0Ts9P312bmNMuH5IiFWIR4RUtLb70Ke5s=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_cloudbuild",
    importpath = "cloud.google.com/go/cloudbuild",
    sum = "h1:TAAmCmAlOJ4uNBu6zwAjwhyl/7fLHHxIEazVhr3QBbQ=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_clouddms",
    importpath = "cloud.google.com/go/clouddms",
    sum = "h1:UhzHIlgFfMr6luVYVNydw/pl9/U5kgtjCMJHnSvoVws=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_cloudtasks",
    importpath = "cloud.google.com/go/cloudtasks",
    sum = "h1:faUiUgXjW8yVZ7XMnKHKm1WE4OldPBUWWfIRN/3z1dc=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_contactcenterinsights",
    importpath = "cloud.google.com/go/contactcenterinsights",
    sum = "h1:tTQLI/ZvguUf9Hv+36BkG2+/PeC8Ol1q4pBW+tgCx0A=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_container",
    importpath = "cloud.google.com/go/container",
    sum = "h1:nbEK/59GyDRKKlo1SqpohY1TK8LmJ2XNcvS9Gyom2A0=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_containeranalysis",
    importpath = "cloud.google.com/go/containeranalysis",
    sum = "h1:2824iym832ljKdVpCBnpqm5K94YT/uHTVhNF+dRTXPI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_datacatalog",
    importpath = "cloud.google.com/go/datacatalog",
    sum = "h1:6kZ4RIOW/uT7QWC5SfPfq/G8sYzr/v+UOmOAxy4Z1TE=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_dataflow",
    importpath = "cloud.google.com/go/dataflow",
    sum = "h1:CW3541Fm7KPTyZjJdnX6NtaGXYFn5XbFC5UcjgALKvU=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_dataform",
    importpath = "cloud.google.com/go/dataform",
    sum = "h1:vLwowLF2ZB5J5gqiZCzv076lDI/Rd7zYQQFu5XO1PSg=",
    version = "v0.5.0",
)

go_repository(
    name = "com_google_cloud_go_datafusion",
    importpath = "cloud.google.com/go/datafusion",
    sum = "h1:j5m2hjWovTZDTQak4MJeXAR9yN7O+zMfULnjGw/OOLg=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_datalabeling",
    importpath = "cloud.google.com/go/datalabeling",
    sum = "h1:dp8jOF21n/7jwgo/uuA0RN8hvLcKO4q6s/yvwevs2ZM=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_dataplex",
    importpath = "cloud.google.com/go/dataplex",
    sum = "h1:cNxeA2DiWliQGi21kPRqnVeQ5xFhNoEjPRt1400Pm8Y=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_dataproc",
    importpath = "cloud.google.com/go/dataproc",
    sum = "h1:gVOqNmElfa6n/ccG/QDlfurMWwrK3ezvy2b2eDoCmS0=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_dataqna",
    importpath = "cloud.google.com/go/dataqna",
    sum = "h1:gx9jr41ytcA3dXkbbd409euEaWtofCVXYBvJz3iYm18=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_datastream",
    importpath = "cloud.google.com/go/datastream",
    sum = "h1:PgIgbhedBtYBU6POGXFMn2uSl9vpqubc3ewTNdcU8Mk=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_deploy",
    importpath = "cloud.google.com/go/deploy",
    sum = "h1:kI6dxt8Ml0is/x7YZjLveTvR7YPzXAUD/8wQZ2nH5zA=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_dialogflow",
    importpath = "cloud.google.com/go/dialogflow",
    sum = "h1:HYHVOkoxQ9bSfNIelSZYNAtUi4CeSrCnROyOsbOqPq8=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_dlp",
    importpath = "cloud.google.com/go/dlp",
    sum = "h1:9I4BYeJSVKoSKgjr70fLdRDumqcUeVmHV4fd5f9LR6Y=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_documentai",
    importpath = "cloud.google.com/go/documentai",
    sum = "h1:jfq09Fdjtnpnmt/MLyf6A3DM3ynb8B2na0K+vSXvpFM=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_domains",
    importpath = "cloud.google.com/go/domains",
    sum = "h1:pu3JIgC1rswIqi5romW0JgNO6CTUydLYX8zyjiAvO1c=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_edgecontainer",
    importpath = "cloud.google.com/go/edgecontainer",
    sum = "h1:hd6J2n5dBBRuAqnNUEsKWrp6XNPKsaxwwIyzOPZTokk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_google_cloud_go_essentialcontacts",
    importpath = "cloud.google.com/go/essentialcontacts",
    sum = "h1:b6csrQXCHKQmfo9h3dG/pHyoEh+fQG1Yg78a53LAviY=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_eventarc",
    importpath = "cloud.google.com/go/eventarc",
    sum = "h1:AgCqrmMMIcel5WWKkzz5EkCUKC3Rl5LNMMYsS+LvsI0=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_filestore",
    importpath = "cloud.google.com/go/filestore",
    sum = "h1:yjKOpzvqtDmL5AXbKttLc8j0hL20kuC1qPdy5HPcxp0=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_functions",
    importpath = "cloud.google.com/go/functions",
    sum = "h1:35tgv1fQOtvKqH/uxJMzX3w6usneJ0zXpsFr9KAVhNE=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_gaming",
    importpath = "cloud.google.com/go/gaming",
    sum = "h1:97OAEQtDazAJD7yh/kvQdSCQuTKdR0O+qWAJBZJ4xiA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_gkebackup",
    importpath = "cloud.google.com/go/gkebackup",
    sum = "h1:4K+jiv4ocqt1niN8q5Imd8imRoXBHTrdnJVt/uFFxF4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_gkeconnect",
    importpath = "cloud.google.com/go/gkeconnect",
    sum = "h1:zAcvDa04tTnGdu6TEZewaLN2tdMtUOJJ7fEceULjguA=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_gkehub",
    importpath = "cloud.google.com/go/gkehub",
    sum = "h1:JTcTaYQRGsVm+qkah7WzHb6e9sf1C0laYdRPn9aN+vg=",
    version = "v0.10.0",
)

go_repository(
    name = "com_google_cloud_go_gkemulticloud",
    importpath = "cloud.google.com/go/gkemulticloud",
    sum = "h1:8F1NhJj8ucNj7lK51UZMtAjSWTgP1zO18XF6vkfiPPU=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_gsuiteaddons",
    importpath = "cloud.google.com/go/gsuiteaddons",
    sum = "h1:TGT2oGmO5q3VH6SjcrlgPUWI0njhYv4kywLm6jag0to=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_iam",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:k4MuwOsS7zGJJ+QfZ5vBK8SgHBAvYN/23BWsiihJ1vs=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_iap",
    importpath = "cloud.google.com/go/iap",
    sum = "h1:BGEXovwejOCt1zDk8hXq0bOhhRu9haXKWXXXp2B4wBM=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_ids",
    importpath = "cloud.google.com/go/ids",
    sum = "h1:LncHK4HHucb5Du310X8XH9/ICtMwZ2PCfK0ScjWiJoY=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_iot",
    importpath = "cloud.google.com/go/iot",
    sum = "h1:Y9+oZT9jD4GUZzORXTU45XsnQrhxmDT+TFbPil6pRVQ=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_kms",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:OWRZzrPmOZUzurjI2FBGtgY2mB1WaJkqhw6oIwSj0Yg=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_language",
    importpath = "cloud.google.com/go/language",
    sum = "h1:3Wa+IUMamL4JH3Zd3cDZUHpwyqplTACt6UZKRD2eCL4=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_lifesciences",
    importpath = "cloud.google.com/go/lifesciences",
    sum = "h1:tIqhivE2LMVYkX0BLgG7xL64oNpDaFFI7teunglt1tI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_longrunning",
    importpath = "cloud.google.com/go/longrunning",
    sum = "h1:NjljC+FYPV3uh5/OwWT6pVU+doBqMg2x/rZlE+CamDs=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_managedidentities",
    importpath = "cloud.google.com/go/managedidentities",
    sum = "h1:3Kdajn6X25yWQFhFCErmKSYTSvkEd3chJROny//F1A0=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_mediatranslation",
    importpath = "cloud.google.com/go/mediatranslation",
    sum = "h1:qAJzpxmEX+SeND10Y/4868L5wfZpo4Y3BIEnIieP4dk=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_memcache",
    importpath = "cloud.google.com/go/memcache",
    sum = "h1:yLxUzJkZVSH2kPaHut7k+7sbIBFpvSh1LW9qjM2JDjA=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_metastore",
    importpath = "cloud.google.com/go/metastore",
    sum = "h1:3KcShzqWdqxrDEXIBWpYJpOOrgpDj+HlBi07Grot49Y=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_monitoring",
    importpath = "cloud.google.com/go/monitoring",
    sum = "h1:c9riaGSPQ4dUKWB+M1Fl0N+iLxstMbCktdEwYSPGDvA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_networkconnectivity",
    importpath = "cloud.google.com/go/networkconnectivity",
    sum = "h1:BVdIKaI68bihnXGdCVL89Jsg9kq2kg+II30fjVqo62E=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_networkmanagement",
    importpath = "cloud.google.com/go/networkmanagement",
    sum = "h1:mDHA3CDW00imTvC5RW6aMGsD1bH+FtKwZm/52BxaiMg=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_networksecurity",
    importpath = "cloud.google.com/go/networksecurity",
    sum = "h1:qDEX/3sipg9dS5JYsAY+YvgTjPR63cozzAWop8oZS94=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_notebooks",
    importpath = "cloud.google.com/go/notebooks",
    sum = "h1:AC8RPjNvel3ExgXjO1YOAz+teg9+j+89TNxa7pIZfww=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_optimization",
    importpath = "cloud.google.com/go/optimization",
    sum = "h1:7PxOq9VTT7TMib/6dMoWpMvWS2E4dJEvtYzjvBreaec=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_orchestration",
    importpath = "cloud.google.com/go/orchestration",
    sum = "h1:39d6tqvNjd/wsSub1Bn4cEmrYcet5Ur6xpaN+SxOxtY=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_orgpolicy",
    importpath = "cloud.google.com/go/orgpolicy",
    sum = "h1:erF5PHqDZb6FeFrUHiYj2JK2BMhsk8CyAg4V4amJ3rE=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_osconfig",
    importpath = "cloud.google.com/go/osconfig",
    sum = "h1:NO0RouqCOM7M2S85Eal6urMSSipWwHU8evzwS+siqUI=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_oslogin",
    importpath = "cloud.google.com/go/oslogin",
    sum = "h1:pKGDPfeZHDybtw48WsnVLjoIPMi9Kw62kUE5TXCLCN4=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_phishingprotection",
    importpath = "cloud.google.com/go/phishingprotection",
    sum = "h1:OrwHLSRSZyaiOt3tnY33dsKSedxbMzsXvqB21okItNQ=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_policytroubleshooter",
    importpath = "cloud.google.com/go/policytroubleshooter",
    sum = "h1:NQklJuOUoz1BPP+Epjw81COx7IISWslkZubz/1i0UN8=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_privatecatalog",
    importpath = "cloud.google.com/go/privatecatalog",
    sum = "h1:Vz86uiHCtNGm1DeC32HeG2VXmOq5JRYA3VRPf8ZEcSg=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise_v2",
    importpath = "cloud.google.com/go/recaptchaenterprise/v2",
    sum = "h1:UqzFfb/WvhwXGDF1eQtdHLrmni+iByZXY4h3w9Kdyv8=",
    version = "v2.5.0",
)

go_repository(
    name = "com_google_cloud_go_recommendationengine",
    importpath = "cloud.google.com/go/recommendationengine",
    sum = "h1:6w+WxPf2LmUEqX0YyvfCoYb8aBYOcbIV25Vg6R0FLGw=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_recommender",
    importpath = "cloud.google.com/go/recommender",
    sum = "h1:9kMZQGeYfcOD/RtZfcNKGKtoex3DdoB4zRgYU/WaIwE=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_redis",
    importpath = "cloud.google.com/go/redis",
    sum = "h1:/zTwwBKIAD2DEWTrXZp8WD9yD/gntReF/HkPssVYd0U=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_resourcemanager",
    importpath = "cloud.google.com/go/resourcemanager",
    sum = "h1:NDao6CHMwEZIaNsdWy+tuvHaavNeGP06o1tgrR0kLvU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_resourcesettings",
    importpath = "cloud.google.com/go/resourcesettings",
    sum = "h1:eTzOwB13WrfF0kuzG2ZXCfB3TLunSHBur4s+HFU6uSM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_retail",
    importpath = "cloud.google.com/go/retail",
    sum = "h1:N9fa//ecFUOEPsW/6mJHfcapPV0wBSwIUwpVZB7MQ3o=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_run",
    importpath = "cloud.google.com/go/run",
    sum = "h1:AWPuzU7Xtaj3Jf+QarDWIs6AJ5hM1VFQ+F6Q+VZ6OT4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_scheduler",
    importpath = "cloud.google.com/go/scheduler",
    sum = "h1:K/mxOewgHGeKuATUJNGylT75Mhtjmx1TOkKukATqMT8=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_secretmanager",
    importpath = "cloud.google.com/go/secretmanager",
    sum = "h1:xE6uXljAC1kCR8iadt9+/blg1fvSbmenlsDN4fT9gqw=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_security",
    importpath = "cloud.google.com/go/security",
    sum = "h1:KSKzzJMyUoMRQzcz7azIgqAUqxo7rmQ5rYvimMhikqg=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_securitycenter",
    importpath = "cloud.google.com/go/securitycenter",
    sum = "h1:QTVtk/Reqnx2bVIZtJKm1+mpfmwRwymmNvlaFez7fQY=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_servicecontrol",
    importpath = "cloud.google.com/go/servicecontrol",
    sum = "h1:ImIzbOu6y4jL6ob65I++QzvqgFaoAKgHOG+RU9/c4y8=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_servicedirectory",
    importpath = "cloud.google.com/go/servicedirectory",
    sum = "h1:f7M8IMcVzO3T425AqlZbP3yLzeipsBHtRza8vVFYMhQ=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_servicemanagement",
    importpath = "cloud.google.com/go/servicemanagement",
    sum = "h1:TpkCO5M7dhKSy1bKUD9o/sSEW/U1Gtx7opA1fsiMx0c=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_serviceusage",
    importpath = "cloud.google.com/go/serviceusage",
    sum = "h1:b0EwJxPJLpavSljMQh0RcdHsUrr5DQ+Nelt/3BAs5ro=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_shell",
    importpath = "cloud.google.com/go/shell",
    sum = "h1:b1LFhFBgKsG252inyhtmsUUZwchqSz3WTvAIf3JFo4g=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_speech",
    importpath = "cloud.google.com/go/speech",
    sum = "h1:yK0ocnFH4Wsf0cMdUyndJQ/hPv02oTJOxzi6AgpBy4s=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_storagetransfer",
    importpath = "cloud.google.com/go/storagetransfer",
    sum = "h1:fUe3OydbbvHcAYp07xY+2UpH4AermGbmnm7qdEj3tGE=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_talent",
    importpath = "cloud.google.com/go/talent",
    sum = "h1:MrekAGxLqAeAol4Sc0allOVqUGO8j+Iim8NMvpiD7tM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_texttospeech",
    importpath = "cloud.google.com/go/texttospeech",
    sum = "h1:ccPiHgTewxgyAeCWgQWvZvrLmbfQSFABTMAfrSPLPyY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_tpu",
    importpath = "cloud.google.com/go/tpu",
    sum = "h1:ztIdKoma1Xob2qm6QwNh4Xi9/e7N3IfvtwG5AcNsj1g=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_trace",
    importpath = "cloud.google.com/go/trace",
    sum = "h1:qO9eLn2esajC9sxpqp1YKX37nXC3L4BfGnPS0Cx9dYo=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_translate",
    importpath = "cloud.google.com/go/translate",
    sum = "h1:AOYOH3MspzJ/bH1YXzB+xTE8fMpn3mwhLjugwGXvMPI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_video",
    importpath = "cloud.google.com/go/video",
    sum = "h1:ttlvO4J5c1VGq6FkHqWPD/aH6PfdxujHt+muTJlW1Zk=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_videointelligence",
    importpath = "cloud.google.com/go/videointelligence",
    sum = "h1:RPFgVVXbI2b5vnrciZjtsUgpNKVtHO/WIyXUhEfuMhA=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_vision_v2",
    importpath = "cloud.google.com/go/vision/v2",
    sum = "h1:TQHxRqvLMi19azwm3qYuDbEzZWmiKJNTpGbkNsfRCik=",
    version = "v2.5.0",
)

go_repository(
    name = "com_google_cloud_go_vmmigration",
    importpath = "cloud.google.com/go/vmmigration",
    sum = "h1:A2Tl2ZmwMRpvEmhV2ibISY85fmQR+Y5w9a0PlRz5P3s=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_vpcaccess",
    importpath = "cloud.google.com/go/vpcaccess",
    sum = "h1:woHXXtnW8b9gLFdWO9HLPalAddBQ9V4LT+1vjKwR3W8=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_webrisk",
    importpath = "cloud.google.com/go/webrisk",
    sum = "h1:ypSnpGlJnZSXbN9a13PDmAYvVekBLnGKxQ3Q9SMwnYY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_websecurityscanner",
    importpath = "cloud.google.com/go/websecurityscanner",
    sum = "h1:y7yIFg/h/mO+5Y5aCOtVAnpGUOgqCH5rXQ2Oc8Oq2+g=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_workflows",
    importpath = "cloud.google.com/go/workflows",
    sum = "h1:7Chpin9p50NTU8Tb7qk+I11U/IwVXmDhEoSsdccvInE=",
    version = "v1.9.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
    version = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:+kGHl1aib/qcwaRi1CbqBZ1rk19r85MNUf8HaBghugY=",
    version = "v0.0.0-20200526211855-cb27e3aa2013",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/grpc",
    sum = "h1:E1eGv1FTqoLIdnBCZufiSHgKjlqG6fKFf6pPWtMTh8U=",
    version = "v1.51.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:SnqbnDw1V7RiZcXPx5MEeqPv2s79L9i7BJUlG/+RurQ=",
    version = "v1.27.1",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:6zppjxzCulZykYSLyVDYbneBfbaBIQPYMevg0bEwv2s=",
    version = "v0.6.0-dev.0.20220419223038-86c51ed26bb4",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:PxfKdU9lEEDYjdIzOtC4qFWgkU2rGHdKlKowJSMN9h0=",
    version = "v0.0.0-20220722155237-a158d28d115b",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:TzXSXBo42m9gQenoE3b9BGiEpg5IG2JkU5FkPIawgtw=",
    version = "v0.0.0-20200107190931-bf48bf16ab8d",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:v4INt8xihDGvnrfjMDVXGxw9wrfxYyCjk0KbXjhR55s=",
    version = "v0.0.0-20220722155257-8c9f86f7a55f",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:JGgROgKl9N8DuW20oFS5gxc+lE67/N3FcwmBPMe7ArY=",
    version = "v0.0.0-20210927222741-03fcf44c2211",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:BrVqGRd7+k1DiOgtnFvAkoQEWQvBc25ouMJM6429SFg=",
    version = "v0.4.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:VveCTK38A2rkS8ZqFY25HIDFscX5X9OoEhJd3quQmXU=",
    version = "v0.1.12",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
    version = "v0.0.0-20190523083050-ea95bdfd59fc",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
    version = "v0.0.0-20190812154241-14fe0d1b01d4",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:VklqNMn3ovrHsnt90PveolxSbWFaJdECFbxSq0Mqo2M=",
    version = "v0.0.0-20190308221718-c2843e01d9a2",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
    version = "v0.0.0-20190121172915-509febef88a4",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
    version = "v0.0.0-20190313153728-d0100b6bd8b3",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:8gQV6CLnAEikrhgkHFbMAEhagSSnXWGV915qUMm9mrU=",
    version = "v0.0.0-20190423024810-112230192c58",
)

go_rules_dependencies()

go_register_toolchains(version = "1.19.3")

gazelle_dependencies(go_repository_default_config = "//:WORKSPACE")

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_google_googleapis",
    branch = "master",
    remote = "https://github.com/googleapis/googleapis.git",
)

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
)

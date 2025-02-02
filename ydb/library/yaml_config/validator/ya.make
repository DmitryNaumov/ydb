LIBRARY()

SRCS(
    validator_builder.h
    validator_builder.cpp
    validator.h
    validator.cpp
    validator_checks.h
    validator_checks.cpp
)

PEERDIR(
    library/cpp/yaml/fyamlcpp
)

GENERATE_ENUM_SERIALIZATION(validator_builder.h)

END()

RECURSE_FOR_TESTS(
    ut
)

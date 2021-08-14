from python_codegen.core.hello_world import print_message


def main() -> int:
    print_message(msg="Hello World!")
    return 0


if __name__ == "__main__":
    exit(main())

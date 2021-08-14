from python_codegen.core import print_message


def test_print_message():
    # Execution
    result = print_message(

        'Hello World',

    )

    # Assertation

    assert result is None

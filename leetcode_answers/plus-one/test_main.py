import pytest

from main import Solution


@pytest.mark.parametrize(
    "test_input,expected",
    [
        ([1, 2, 3], [1, 2, 4]),
        ([4, 3, 2, 1], [4, 3, 2, 2]),
        ([9], [1, 0]),
        ([9, 9], [1, 0, 0]),
        ([8, 9, 9, 9], [9, 0, 0, 0]),
    ],
)
def test(test_input, expected):
    assert Solution().plusOne(test_input) == expected

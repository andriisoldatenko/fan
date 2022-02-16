from main import Solution
import pytest


@pytest.mark.parametrize(
    "test_input,expected",
    [
        ([2, 2, 1], 1),
        ([4, 1, 2, 1, 2], 4),
        ([1], 1)
    ]
)
def test_solution(test_input, expected):
    s = Solution()
    assert s.singleNumber(test_input) == expected

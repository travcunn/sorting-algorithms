import random
import unittest

import bogo
import bubble
import insertion
import merge


def generate_list(size):
    return [random.randint(0, 1000) for r in xrange(size)]

def sorted_list(list):
    return sorted(list)


class BubbleSortTest(unittest.TestCase):
    def test_sort(self):
        list = generate_list(1000)
        x = bubble.sort(list)
        self.assertTrue(x == sorted_list(list))


class MergeSortTest(unittest.TestCase):
    def test_sort(self):
        list = generate_list(1000)
        x = merge.sort(list)
        self.assertTrue(x == sorted_list(list))


class InsertionSortTest(unittest.TestCase):
    def test_sort(self):
        list = generate_list(1000)
        x = insertion.sort(list)
        self.assertTrue(x == sorted_list(list))


class BogoSortTest(unittest.TestCase):
    def test_sort(self):
        list = generate_list(10)
        x = bogo.sort(list)
        self.assertTrue(x == sorted_list(x))


if __name__ == '__main__':
    unittest.main()


class Solution(object):
    def palindrome(self, num):
        if num < 10:
            return True
        str_num = str(num)      # space: O(n)
        len_num = len(str_num)
        res = 0
        l = 0
        r = len_num-1
        while r > l:            # time: O(n)
            if str_num[l] != str_num[r]:
                return False
            l += 1
            r -= 1
        return True

    def print_is_palindrome(self, num):
        ret = self.palindrome(num)
        print(f"Is {num} palindrome {ret}")


nums = [555, 1234, 1221, 12521]
sol = Solution()

for num in nums:
    sol.print_is_palindrome(num)


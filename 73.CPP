//初版，的空间复杂度为m+n;并未实现题目要求的原地排序
class Solution {
public:
    void setZeroes(vector<vector<int>> &matrix) {
        int m = matrix.size();
        int n = matrix[0].size();
        vector<int> hline(m,1);
        vector<int> sline(n,1);
        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 0) {
                    hline[i] = 0;
                    sline[j] = 0;
                }
            }

        for (int i = 0; i < m; i++)
            for (int j = 0; j < n; j++) {
                if ((hline[i] != 0) && (sline[j] != 0))
                    continue;
                matrix[i][j] = 0;
            }   
    }
};
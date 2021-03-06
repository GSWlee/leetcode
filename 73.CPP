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


//更新后的函数，实现原地排序
void setZeroes(vector<vector<int>> &matrix) {
    bool hfirst = false;
    bool lfirst = false;
    for (int i = 0; i < matrix.size(); i++)
        if (matrix[i][0] == 0)
            lfirst = true;
    for (int i = 0; i < matrix[0].size(); i++)
        if (matrix[0][i] == 0)
            hfirst = true;

    for (int i = 1; i < matrix.size(); i++)
        for (int j = 1; j < matrix[0].size(); j++) {
            if ((matrix[i][j] == 0)) {
                matrix[0][j] = 0;
                matrix[i][0] = 0;
            }
        }
    for (int i = 1; i < matrix.size(); i++)
        for (int j = 1; j < matrix[0].size(); j++) {
            if ((matrix[i][0] == 0) || (matrix[0][j] == 0)) {
                matrix[i][j] = 0;
            }
        }
    if (hfirst)
        for (int i = 0; i < matrix[0].size(); i++)
            matrix[0][i] = 0;
    if (lfirst)
        for (int i = 0; i < matrix.size(); i++)
            matrix[i][0] = 0;

}
class Solution {
public:
    int minCost(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        vector<bitset<1024>> prev(n),curr(n);

        for(int i=0;i<m;++i){
            for(int j=0;j<n;++j){
                curr[j].reset();
                int v = grid[i][j];
                if(i==0 && j==0){
                    curr[j][v] = 1;
                }

                if(i>0){
                    for(int q=0;q<1024;++q){
                        if(prev[j][q]){
                            curr[j][q^v] = 1;
                        }
                    }
                }
                if(j>0){
                    for(int q=0;q<1024;++q){
                        if(curr[j-1][q]){
                            curr[j][q^v] = 1;
                        }
                    }
                }
                
            }swap(prev,curr);

            
        }
        for(int i=0;i<1024;++i){
                if(prev[n-1][i])return i;
            }
            return -1;
    }
};
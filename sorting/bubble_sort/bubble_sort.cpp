#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void bubbleSort(vector<int>& arr){
    int n = arr.size();
    for(int i=0;i<n;i++){
        bool swapped = false;
        for(int j=0;j<n-i-1;j++){
            if(arr[j]>arr[j+1]){
                swapped = true;
                swap(arr[j],arr[j+1]);
            }
        }
        if(!swapped){
            return;
        }
    }
    return;
}

int main(){
    vector<int> arr = {2,4,3,5,2,1,5,6};
    bubbleSort(arr);
    for(int num : arr){
        cout << num << " ";
    }
    cout << endl;
    return 0;
}

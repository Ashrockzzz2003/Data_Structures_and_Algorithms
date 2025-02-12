#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void SelectionSort(vector<int>& arr){
    int n = arr.size();
    for(int i=0;i<n;i++){
        int min_idx = i;
        for(int j=i+1;j<n;j++){
            if(arr[j]<arr[min_idx]){
                min_idx = j;
            }
        }
        if(min_idx!=i){
            swap(arr[i],arr[min_idx]);
        }
    }
}

int main(){
    vector<int> arr = {6,4,3,2,5,6,43,67,3,5,5};
    SelectionSort(arr);
    for(int num : arr){
        cout << num << " ";
    }
    cout << endl;
}
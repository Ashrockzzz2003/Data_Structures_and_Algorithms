#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void InsertionSort(vector<int>& arr){
    int n = arr.size();
    for(int i=1;i<n;i++){
        int key = arr[i];
        int j = i-1;
        while(j>=0 && key < arr[j]){
            arr[j+1] = arr[j];
            j-=1;
        }
        arr[j+1] = key;
    }
}

int main(){
    vector<int> arr = {6,4,3,2,5,6,43,67,3,5,5};
    InsertionSort(arr);
    for(int num : arr){
        cout << num << " ";
    }
    cout << endl;
}

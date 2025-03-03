#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void merge(vector<int>& arr , int l , int r , int m){
    vector<int> temp;
    int left = l;
    int right = m+1;
    while(left<= m && right <= r){
        if(arr[left]<arr[right]){
            temp.push_back(arr[left]);
            left++;
        }
        else{
            temp.push_back(arr[right]);
            right++;
        }
    }

    while(left<=m){
        temp.push_back(arr[left]);
        left++;
    }
    while(right<=r){
        temp.push_back(arr[right]);
        right++;
    }

    for(int i=l;i<=r;i++){
        arr[i] = temp[i-l];
    }

}

void mergeSort(vector<int>& arr , int l , int r){
    if (l<r){
        int mid = (l+r)/2;
        mergeSort(arr,l,mid);
        mergeSort(arr,mid+1,r);
        merge(arr,l,r,mid);
    }
}

int main(){
    vector<int> arr = {2,4,4,2,24,21,34,44,1,3,3,45,6};
    mergeSort(arr,0,arr.size()-1);
    for(int num : arr){
        cout << num << " ";
    }
    cout << endl;
}

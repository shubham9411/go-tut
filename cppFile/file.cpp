#include <iostream>
#include <fstream>
using namespace std;

struct Data
{
  int data = 2;
};

int main()
{
  string FILE_NAME = "filename.txt";
  Data datum;
  datum.data = 100;
  // Create and open a text file
  fstream file;
  // file.open("filename.txt", ios::out|ios::binary);

  fstream MyFile("filename.txt", ios::out | ios::binary);

  // Write to the file
  // MyFile << "Files can be tricky, but it is fun enough!";
  // MyFile << datum;
  MyFile.write((char *)&datum, sizeof(datum));
  // Close the file
  MyFile.close();

  file.open(FILE_NAME, ios::in | ios::binary);
  if (!file)
  {
    cout << "Error in opening file...\n";
    return -1;
  }

  if (file.read((char *)&datum, sizeof(datum)))
  {
    cout << endl
         << endl;
    cout << "Data extracted from file..\n";
    // print the object
    cout<<datum.data<<endl;
  }
  else
  {
    cout << "Error in reading data from file...\n";
    return -1;
  }

  file.close();
}
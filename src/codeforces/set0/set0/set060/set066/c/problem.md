Recently, on a programming lesson little Petya showed how quickly he can create files and folders on the computer. But he got soon fed up with this activity, and he decided to do a much more useful thing. He decided to calculate what folder contains most subfolders (including nested folders, nested folders of nested folders, and so on) and what folder contains most files (including the files in the subfolders).

More formally, the subfolders of the folder are all its directly nested folders and the subfolders of these nested folders. The given folder is not considered the subfolder of itself. A file is regarded as lying in a folder, if and only if it either lies directly in this folder, or lies in some subfolder of the folder.

For a better understanding of how to count subfolders and files for calculating the answer, see notes and answers to the samples.

You are given a few files that Petya has managed to create. The path to each file looks as follows:

diskName:\folder1\folder2\...\ foldern\fileName

diskName is single capital letter from the set {C,D,E,F,G}.
folder1, ..., foldern are folder names. Each folder name is nonempty sequence of lowercase Latin letters and digits from 0 to 9. (n ≥ 1)
fileName is a file name in the form of name.extension, where the name and the extension are nonempty sequences of lowercase Latin letters and digits from 0 to 9.
It is also known that there is no file whose path looks like diskName:\fileName. That is, each file is stored in some folder, but there are no files directly in the root. Also let us assume that the disk root is not a folder.

Help Petya to find the largest number of subfolders, which can be in some folder, and the largest number of files that can be in some folder, counting all its subfolders.
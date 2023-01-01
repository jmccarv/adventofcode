typedef struct pairs pair_t;
typedef struct pairs {
    int nr;
    unsigned char min1[1000];
    unsigned char max1[1000];
    unsigned char min2[1000];
    unsigned char max2[1000];
};

const pair_t assn = {
1000,
{18,9,7,82,17,13,46,34,4,9,66,1,24,3,20,93,1,95,5,93,2,11,2,54,7,28,8,38,97,20,57,34,7,40,50,25,13,47,49,11,85,61,19,10,90,49,44,34,68,2,19,12,95,27,86,39,42,23,3,3,76,26,27,33,64,3,52,90,16,30,51,44,3,51,5,59,84,23,43,67,10,33,20,33,32,7,28,13,19,24,90,5,10,25,13,3,4,15,4,22,22,29,11,6,47,2,6,24,10,58,44,4,99,4,86,50,4,2,4,12,46,17,52,43,26,67,38,94,7,3,82,45,2,26,81,1,3,84,48,32,30,3,49,47,2,41,67,10,14,15,35,48,13,7,3,17,24,18,5,21,16,1,2,4,73,4,18,74,2,3,72,6,3,7,7,10,21,54,9,16,5,39,57,42,53,25,12,27,14,41,27,2,53,59,21,2,14,9,83,3,17,15,13,60,47,65,4,97,18,24,38,40,23,19,9,96,48,60,30,13,9,30,56,1,3,32,31,61,20,1,1,8,10,15,81,59,6,17,34,47,22,74,47,95,1,86,32,77,85,39,61,58,82,2,17,31,81,52,59,67,17,22,3,76,63,14,67,29,2,3,3,6,67,84,23,21,24,19,29,15,12,11,66,77,63,5,47,52,53,27,26,27,17,33,79,67,5,20,54,28,9,6,37,52,35,21,30,12,62,34,20,7,88,2,4,38,85,10,30,38,83,40,24,24,56,32,3,11,48,17,31,4,76,9,61,35,56,58,65,47,7,92,11,32,18,68,62,21,13,7,7,1,9,31,14,56,7,13,24,13,28,15,7,23,58,76,30,80,21,3,4,56,2,35,64,10,1,97,50,12,89,14,81,20,19,35,32,39,47,36,4,33,1,17,83,12,68,15,55,9,11,23,18,4,42,7,17,6,13,52,38,30,7,94,18,7,33,56,11,3,69,67,45,1,7,9,64,28,2,2,84,4,52,43,56,12,68,31,1,47,98,8,9,15,28,16,15,12,55,51,37,9,24,29,43,47,60,48,48,18,4,15,5,77,23,3,5,5,6,42,43,9,88,44,21,59,17,49,36,30,42,16,3,18,14,25,36,17,3,36,1,54,26,15,31,22,74,24,24,25,24,17,19,4,54,61,96,50,66,8,5,3,74,23,9,18,67,13,36,70,20,25,25,21,33,90,39,5,35,11,56,26,2,66,39,2,20,30,14,38,53,9,3,6,96,30,80,6,11,35,41,58,53,2,5,8,84,47,64,93,70,67,31,27,50,3,90,59,14,4,10,34,13,15,7,31,41,1,50,16,8,61,3,9,19,15,1,31,2,20,63,4,3,4,20,12,32,17,31,72,70,41,82,6,6,45,10,48,28,2,41,2,1,76,1,5,24,10,13,98,1,18,4,72,7,1,39,33,76,4,6,1,2,37,5,39,71,33,10,9,11,7,7,13,16,11,10,1,25,78,23,31,50,2,64,8,34,34,22,74,85,9,29,15,82,66,3,17,93,86,15,22,36,53,87,32,15,10,25,4,35,3,6,46,19,38,13,18,2,2,32,48,57,15,98,35,57,2,7,27,30,90,3,37,35,22,4,67,11,70,96,11,48,20,44,27,98,40,3,51,4,4,2,50,64,2,14,8,35,5,58,11,27,3,37,45,21,3,52,67,31,69,22,32,4,4,87,42,43,48,30,11,61,92,94,41,77,7,94,18,47,29,22,9,43,11,9,22,42,1,2,14,17,8,94,9,2,1,25,1,13,97,55,84,8,17,54,62,66,37,13,8,14,49,24,84,19,27,17,23,61,2,24,14,81,23,42,60,4,3,3,6,26,8,1,26,90,60,57,18,19,79,41,25,4,15,46,52,6,15,22,61,75,67,5,2,13,2,75,55,39,62,44,7,31,40,6,4,8,11,56,1,52,2,1,3,1,15,65,21,54,24,32,4,1,6,48,5,22,4,28,3,6,26,3,88,77,19,4,10,1,22,35,1,20,2,1,14,7,16,27,86,31,25,4,51,18,37,80,2,54,11,11,65,22,8,61,27,90,9,3,73,13,71,44,56,47,50,6,49,48,98,18,50,7,79,4,79,20,26,15,2,5,2,1,84,18,50,82,17,82,20,7,2,2,57,70,37,51,18,76,15,3,60,11,8,31,92,14,31,32,31,18,19,73,31,13,2,13,53,8,28,5,96,4,1,33,18,11,36,48,11,67,13,43,68,89,27,31,9,75,19,99,69,9,83,22,41,2,47},
{20,86,8,98,17,21,52,54,91,80,83,5,27,91,81,99,86,99,94,95,87,97,49,59,96,29,86,54,97,88,80,39,73,62,52,79,84,83,77,53,86,78,52,66,92,88,79,57,70,5,74,99,95,74,90,96,49,42,86,24,80,38,39,77,81,4,80,91,56,99,96,82,53,71,80,61,84,78,84,92,98,57,72,98,55,88,76,39,26,99,96,10,92,89,97,97,6,33,42,93,84,34,66,38,93,97,55,90,79,59,55,88,99,40,92,82,77,98,34,52,47,89,66,51,68,99,92,94,93,88,83,74,95,93,82,35,70,90,65,76,86,6,80,88,13,73,92,85,87,23,56,78,99,9,85,42,68,71,33,22,83,15,2,50,97,95,62,80,62,75,74,98,98,73,77,69,92,60,83,17,81,81,87,95,60,25,66,54,30,48,58,12,54,61,80,37,85,38,97,90,97,18,90,91,85,72,44,98,97,49,85,90,91,97,90,98,59,75,99,45,24,61,97,55,99,74,35,90,57,93,94,93,12,30,84,75,78,26,67,53,98,89,59,96,47,87,32,84,85,79,88,59,90,99,96,89,83,80,68,69,87,99,98,94,99,19,68,97,51,83,64,53,69,88,28,96,91,94,91,98,47,86,90,94,81,67,49,93,86,27,94,27,90,78,97,68,63,34,81,99,89,94,40,77,45,66,33,72,84,54,98,72,90,48,94,58,86,78,30,91,85,82,29,70,69,84,98,90,75,95,79,4,91,80,91,86,66,69,90,72,14,97,49,44,74,82,90,93,18,94,10,92,12,56,93,79,96,96,44,77,96,96,43,40,96,76,56,97,49,75,6,91,14,62,93,10,69,98,92,12,92,15,85,84,32,39,64,67,55,67,20,70,98,74,98,36,84,63,56,62,36,81,78,95,75,97,90,85,14,61,42,31,95,95,97,87,33,68,46,96,69,83,74,3,98,31,65,30,70,94,89,33,91,48,58,92,97,33,95,55,98,76,43,73,42,19,28,74,93,76,77,94,49,33,53,95,73,85,68,81,97,16,9,87,99,64,86,80,54,90,45,66,89,71,86,59,95,93,86,73,46,96,30,93,23,68,55,17,91,66,77,70,73,26,48,64,99,26,24,25,81,19,88,7,71,89,98,64,93,83,96,62,94,57,93,29,85,89,75,96,29,79,92,23,74,90,95,11,36,78,85,39,99,68,85,3,32,78,74,40,55,38,99,12,96,32,81,48,54,71,65,90,87,57,65,51,98,99,71,99,72,68,61,79,77,36,91,96,58,4,33,97,99,78,51,46,59,27,83,48,70,85,87,88,21,94,83,32,5,72,64,97,68,98,20,95,33,83,36,78,81,44,82,90,11,45,87,60,73,9,98,89,93,92,70,5,82,25,22,99,83,37,99,73,99,1,89,34,79,48,96,90,99,84,5,92,73,99,16,22,38,85,33,43,43,69,94,33,44,94,85,90,52,54,82,90,88,73,25,77,89,47,93,17,85,77,46,48,98,86,97,63,81,66,94,79,67,85,37,89,76,95,99,50,56,75,48,96,95,95,41,58,86,15,99,93,57,61,96,84,98,92,70,69,51,66,90,77,70,78,97,97,55,82,45,88,99,84,69,69,53,11,99,50,94,90,96,76,84,97,65,91,89,85,52,56,22,31,89,84,75,78,84,36,92,99,88,88,54,49,63,93,87,94,98,69,78,25,95,79,88,73,58,95,66,70,79,33,73,87,20,61,18,95,95,79,98,3,73,77,97,98,80,93,56,94,90,63,67,41,94,73,53,60,91,97,85,79,48,86,61,86,73,93,99,56,89,62,80,71,3,44,97,41,99,93,91,78,59,25,97,87,89,80,81,88,79,71,97,91,96,86,94,73,88,70,15,92,97,63,81,86,99,66,47,88,95,89,99,93,58,85,60,98,92,5,99,67,65,21,64,84,95,84,87,88,79,97,46,67,29,68,94,43,46,90,96,83,26,83,2,22,51,99,45,74,5,98,91,51,63,86,96,98,92,81,86,59,97,99,86,39,13,91,38,32,84,72,94,97,95,73,16,89,72,58,53,60,19,88,60,98,18,82,63,96,96,80,78,85,24,80,74,2,87,85,96,52,82,54,96,66,83,2,97,69,83,96,99,90,99,52,94,81,26,92,80,93,35,85,91,61,20,60,97,84,80,83,58,86,76,61,99,98,95,58,62,48,85,58,75,92,77,83,61,90,90,95,31,97,86,19,99,69,99,91,22,67,8,63},
{19,9,8,98,17,20,45,37,3,5,67,1,23,5,19,15,1,16,93,61,2,11,50,2,8,29,8,38,44,21,58,32,8,39,51,24,12,47,76,30,86,46,20,65,13,87,45,34,52,6,18,12,45,23,88,40,43,5,1,25,8,25,16,34,13,4,42,16,45,1,10,44,2,19,80,60,77,22,14,66,23,33,19,21,55,5,3,12,30,83,66,11,15,39,12,98,5,32,10,92,7,26,11,1,5,97,54,89,10,59,45,4,2,41,15,81,5,97,35,22,46,14,65,51,27,3,93,6,29,1,68,31,37,69,74,36,4,83,23,31,31,1,49,48,3,40,67,9,13,15,57,47,6,7,13,18,23,17,4,22,82,14,2,3,73,4,18,73,3,1,21,6,1,6,6,10,20,99,1,17,3,40,3,94,51,23,11,27,30,40,42,13,18,60,2,36,13,8,99,94,16,16,13,60,86,4,45,9,18,14,37,35,6,29,8,44,37,26,30,12,25,16,56,4,12,32,20,61,56,5,1,3,13,31,82,58,77,17,66,54,23,51,45,11,1,87,32,55,12,39,11,59,5,1,64,88,6,27,60,68,18,11,4,95,64,19,67,29,1,1,4,52,22,5,28,12,90,2,30,14,20,12,39,65,64,6,44,94,25,17,40,12,18,77,96,52,1,10,82,29,91,5,19,14,48,53,29,3,19,81,21,19,12,2,3,14,53,1,31,32,82,82,19,25,56,26,97,10,49,18,72,5,57,88,60,35,57,55,43,46,5,36,11,43,19,31,59,93,17,7,10,2,15,14,15,6,6,8,43,68,28,92,7,23,59,75,29,79,21,23,7,7,15,27,63,11,1,1,31,13,14,14,4,13,32,37,33,68,46,35,3,52,3,16,84,12,1,14,13,61,5,4,77,3,42,5,16,2,14,53,34,23,6,49,17,8,34,55,10,2,63,84,37,4,8,10,65,12,69,4,16,5,52,17,57,8,39,36,96,48,2,75,5,57,28,13,29,65,3,56,61,9,49,28,33,35,61,47,18,18,3,15,8,16,22,63,91,1,6,47,44,10,6,45,85,60,18,45,37,73,38,15,4,18,13,7,28,18,3,35,2,21,26,27,48,21,54,25,25,26,80,11,87,6,72,64,36,64,97,9,6,4,52,5,10,17,16,4,15,2,28,24,17,22,34,69,38,4,36,10,57,38,3,11,81,2,20,5,7,39,54,8,2,11,91,25,4,7,2,56,40,58,25,27,4,7,18,47,70,57,1,96,51,27,51,35,18,35,13,5,10,34,98,72,4,31,41,15,63,47,69,61,2,9,8,20,1,32,4,68,50,13,88,3,21,43,33,82,25,12,69,43,18,1,6,17,9,47,28,9,97,88,2,66,26,6,49,13,17,48,3,35,15,5,8,2,39,7,74,4,16,90,54,38,6,91,19,33,15,9,18,82,6,14,15,22,10,27,32,50,27,91,51,20,65,8,35,14,26,66,25,8,28,11,2,10,6,17,94,39,33,21,37,51,10,31,66,7,25,2,35,1,5,49,19,52,12,4,1,94,31,20,16,16,6,34,58,2,59,27,29,90,3,36,45,1,2,68,7,55,4,12,47,20,45,27,13,40,1,49,2,11,4,49,53,5,53,6,12,96,3,32,26,84,51,44,22,31,88,85,76,77,71,32,91,2,59,41,20,22,29,10,14,4,32,40,32,7,2,34,89,96,58,94,66,11,10,12,42,2,3,15,9,9,7,89,48,3,25,1,12,33,24,59,7,93,54,24,66,42,12,8,39,24,13,2,20,11,18,16,40,1,38,79,62,57,54,7,8,4,4,6,97,41,1,76,16,59,87,25,21,20,40,28,3,15,51,70,6,92,96,85,74,72,6,12,14,1,75,50,76,30,61,7,13,87,96,4,7,10,57,86,52,3,3,4,2,42,62,9,28,51,3,2,1,5,47,4,21,11,29,1,4,26,4,39,50,98,1,82,3,8,34,98,14,1,5,14,9,6,13,54,16,98,4,50,40,25,59,85,55,11,12,98,37,9,83,11,12,9,4,58,12,45,44,57,43,10,20,69,21,85,19,81,8,78,1,1,19,26,15,3,4,3,7,34,18,51,48,18,3,67,8,3,96,70,69,39,50,19,75,6,62,59,11,7,79,3,2,84,32,30,19,20,26,85,12,83,14,22,7,29,4,2,94,6,4,8,86,36,49,91,66,73,60,67,90,26,11,11,76,20,7,69,8,10,21,40,7,46},
{21,87,18,99,77,79,46,53,5,83,83,1,23,90,20,94,87,96,94,94,86,96,78,55,97,93,55,64,96,89,80,35,73,62,51,26,85,84,77,52,86,79,92,66,91,88,80,58,69,85,75,96,96,26,88,97,55,75,93,36,76,97,38,77,63,98,51,96,47,31,51,81,64,67,96,98,85,78,82,68,96,41,71,34,97,96,28,38,71,83,76,96,90,90,97,99,97,33,71,94,16,35,85,7,92,99,86,90,57,60,56,87,96,66,85,92,78,98,35,36,47,97,66,88,68,66,95,70,95,4,83,74,94,88,82,64,70,89,48,76,85,3,79,88,14,74,88,85,87,24,73,85,13,12,51,43,47,71,4,96,97,16,2,4,98,59,99,80,43,18,73,99,3,74,77,34,93,99,3,59,96,81,86,95,61,24,67,88,92,48,59,46,53,60,84,87,86,8,99,98,16,19,89,91,91,66,82,98,47,64,97,56,99,97,89,81,59,59,79,12,94,51,80,93,94,32,34,88,91,92,95,7,55,85,85,75,79,35,67,54,98,80,47,96,37,87,33,98,85,78,61,91,90,99,96,90,82,51,67,91,87,92,98,96,72,22,69,98,3,86,39,98,69,87,82,98,95,98,86,99,66,87,66,94,81,82,48,94,87,72,77,26,91,79,99,68,6,35,82,98,94,94,41,77,99,65,29,84,58,83,93,93,89,23,93,39,86,11,60,82,84,83,74,69,71,85,99,89,75,45,74,89,77,89,91,85,67,58,66,48,6,90,64,45,75,56,91,98,61,8,71,92,70,49,90,57,28,45,44,77,97,94,42,26,96,96,31,96,48,75,96,55,70,34,63,18,70,98,91,59,90,15,99,89,72,40,64,75,48,67,19,69,96,74,97,29,17,63,63,63,12,85,78,97,64,7,16,88,96,62,43,32,94,65,96,88,54,78,46,95,70,84,75,41,99,32,79,29,71,29,83,48,95,47,60,13,69,78,96,55,99,77,25,73,40,20,29,73,93,61,76,24,76,66,54,95,74,47,48,92,5,29,91,76,98,99,93,85,46,80,79,67,89,72,86,91,99,49,81,81,56,95,35,94,45,24,90,98,91,66,77,93,64,96,82,64,73,67,88,63,87,18,88,93,72,76,97,69,99,84,96,69,95,28,93,29,31,13,36,97,93,79,93,70,75,91,95,11,94,90,86,41,99,67,81,60,31,79,82,78,67,38,99,73,95,32,81,48,12,86,66,91,87,56,66,51,28,47,78,93,71,97,52,79,78,89,91,93,13,83,21,52,99,79,35,47,60,17,97,48,99,78,3,88,29,88,84,49,27,71,63,93,96,97,34,83,71,83,32,73,69,65,81,6,11,44,9,60,73,95,99,90,93,75,70,98,61,26,21,96,99,39,99,73,99,32,88,34,79,47,93,91,86,83,99,92,72,93,88,48,37,82,33,44,15,43,93,99,42,96,87,91,52,52,82,89,89,78,76,76,85,47,92,16,81,74,25,49,99,85,98,22,82,54,88,79,67,86,93,3,86,3,57,70,89,76,47,18,96,99,42,78,70,61,97,92,87,97,96,91,97,91,71,38,52,22,3,90,11,69,97,98,54,51,57,68,99,91,73,50,2,23,97,49,95,16,97,46,36,96,89,86,26,84,91,56,84,55,90,85,76,78,95,36,93,4,88,41,54,49,44,10,61,93,94,63,78,24,93,65,94,96,98,98,89,70,80,58,74,87,75,62,18,96,95,91,55,43,90,77,96,83,80,89,55,95,89,61,67,62,94,82,53,46,23,96,86,45,49,87,60,86,72,89,82,74,90,61,79,70,72,81,98,42,98,93,90,93,97,40,82,88,89,55,92,59,83,95,98,92,97,86,97,86,89,66,30,1,97,63,78,96,86,59,30,89,96,91,99,55,73,86,87,98,26,23,99,96,67,20,53,85,96,97,87,88,62,4,23,47,68,3,4,88,53,89,78,98,82,83,82,21,85,98,21,75,21,99,92,55,63,87,96,99,93,82,85,66,81,98,85,37,62,98,45,32,85,72,90,68,95,74,14,72,73,71,52,49,67,87,70,97,30,82,63,95,1,80,81,27,24,81,4,92,84,85,19,89,82,58,97,94,83,67,99,71,83,42,99,90,88,14,95,59,51,92,81,93,13,85,90,61,48,20,85,85,80,98,57,86,9,62,5,97,94,93,63,14,99,59,96,93,67,82,77,68,90,26,36,95,86,64,97,70,99,92,21,42,61,68},
};

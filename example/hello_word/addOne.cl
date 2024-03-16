__kernel void addOne(__global float* data) {
    int i = get_global_id(0);
    printf("id = %03d, data[i]+1 = %.0f\n", i, data[i]+1);
    data[i] = data[i]+1;
}

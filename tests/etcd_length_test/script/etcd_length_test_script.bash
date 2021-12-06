name=`hostname -s`
rm -rf ${name}.etcd
my_int_ip=`hostname -i | awk '{print $1}'`
echo "my internal ip is ${my_int_ip}"
int_ip_1=${my_int_ip}
echo "please enter 2 more internal ip addresses seperated with a space"
read int_ip_2 int_ip_3
echo "first internal ip: ${int_ip_1}, second internal ip: ${int_ip_2}, third internal ip: ${int_ip_3}"
etcd --name="${name}" \
     --initial-advertise-peer-urls http://${my_int_ip}:2380 \
     --listen-peer-urls http://${my_int_ip}:2380 \
     --listen-client-urls http://${my_int_ip}:2379,http://127.0.0.1:2379 \
     --advertise-client-urls http://${my_int_ip}:2379 \
     --initial-cluster-token etcd-cluster-1 \
     --initial-cluster perf-etcd-aot-1=http://${int_ip_1}:2380,perf-etcd-aot-2=http://${int_ip_2}:2380,perf-etcd-aot-3=http://${int_ip_3}:2380 \
     --initial-cluster-state=new \
     --max-txn-ops=1000000 \
	 --max-request-bytes=$(( 1 * 1024 * 1024 * 1024 )) \
     --quota-backend-bytes=21474836480 \
     --auto-compaction-retention=5m \
     -log-level=error

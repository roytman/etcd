name=`hostname -s`
rm -rf ${name}.etcd
int_ip=`hostname -i | awk '{print $1}'`
ext_ip=`curl ifconfig.me`
echo "my external ip is ${ext_ip}"
echo "my internal ip is ${int_ip}"
etcd --name="${name}" \
     --initial-advertise-peer-urls http://${int_ip}:2380 \
     --listen-peer-urls http://${int_ip}:2380 \
     --listen-client-urls http://${int_ip}:2379,http://127.0.0.1:2379 \
     --advertise-client-urls http://${int_ip}:2379 \
     --initial-cluster-token etcd-cluster-1 \
     --initial-cluster perf-etcd-aot-1=http://10.243.64.8:2380,perf-etcd-aot-2=http://10.243.64.9:2380,perf-etcd-aot-3=http://10.243.64.10:2380 \
     --initial-cluster-state=new \
     --max-txn-ops=1000000 \
	 --max-request-bytes=$(( 1 * 1024 * 1024 * 1024 )) \
     --quota-backend-bytes=21474836480 \
     --auto-compaction-retention=5m \
     -log-level=error


// 不要在循环体中使用array_merge! 

<!-- $len = 100;
$modelOrder = new OptionsOrderModel();
date_default_timezone_set('Asia/Shanghai');
$list = $modelOrder->where($condition)->field($field)->order('order_sys_id desc')->limit(10000);
$tempOrderListAll = [[]];
$checkLen = $list->select()->count() / $len;
for ($i = 0; $i < $checkLen; $i++) {
    //$orderListAll = array_merge($orderListAll, $list->limit($i * $len, $len)->select()->toArray());
    $tempOrderListAll[] = $list->limit($i * $len, $len)->select()->toArray();
}
$orderListAll = array_merge(...$tempOrderListAll); -->
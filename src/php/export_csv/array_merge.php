
<!-- https://learnku.com/articles/25111 -->
<!-- 不要在循环体中使用 array_merge () -->

<?php 

/**
 * Class ArrayMerge
 */
class ArrayMerge
{

    /**
     * @param int $times
     * @return array
     */
    public static function eachOne(int $times): array
    {
        $a = [];
        $b = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
        for ($i = 0; $i < $times; $i++) {
            $a = array_merge($a, $b);
        }
        return $a;
    }

    /**
     * @param int $times
     * @return array
     */
    public static function eachTwo(int $times): array
    {
        $a = [[]];
        $b = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
        for ($i = 0; $i < $times; $i++) {
            $a[] = $b;
        }
        return array_merge(...$a);
    }

    /**
     * @param int $times
     * @return array
     */
    public static function eachThree(int $times): array
    {
        $a = [];
        $b = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
        for ($i = 0; $i < $times; $i++) {
            foreach ($b as $item) {
                $a[] = $item;
            }
        }
        return $a;
    }

    /**
     * @param int $times
     * @return array
     */
    public static function eachFour(int $times): array
    {
        $a = [];
        $b = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
        for ($i = 0; $i < $times; $i++) {
            $a = $b + $a;
        }
        return $a;
    }

    /**
     * 转化内存信息
     * @param      $bytes
     * @param bool $binaryPrefix
     * @return string
     */
    public static function getNiceFileSize(int $bytes, $binaryPrefix = true): ?string
    {
        if ($binaryPrefix) {
            $unit = array('B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB');
            if ($bytes === 0) {
                return '0 ' . $unit[0];
            }
            return @round($bytes / (1024 ** ($i = floor(log($bytes, 1024)))),
                    2) . ' ' . ($unit[(int)$i] ?? 'B');
        }

        $unit = array('B', 'KB', 'MB', 'GB', 'TB', 'PB');
        if ($bytes === 0) {
            return '0 ' . $unit[0];
        }
        return @round($bytes / (1000 ** ($i = floor(log($bytes, 1000)))),
                2) . ' ' . ($unit[(int)$i] ?? 'B');
    }
}

ini_set('memory_limit', '4000M');
$timeOne = microtime(true);
$a       = ArrayMerge::eachOne(10000);
echo 'count eachOne Result | ' . count($a) . PHP_EOL;
echo 'memory eachOne Result | ' . ArrayMerge::getNiceFileSize(memory_get_usage(true)) . PHP_EOL;
$timeTwo = microtime(true);
$b       = ArrayMerge::eachTwo(10000);
echo 'count eachTwo Result | ' . count($b) . PHP_EOL;
echo 'memory eachTwo Result | ' . ArrayMerge::getNiceFileSize(memory_get_usage(true)) . PHP_EOL;
$timeThree = microtime(true);
$c         = ArrayMerge::eachThree(10000);
echo 'count eachThree Result | ' . count($c) . PHP_EOL;
echo 'memory eachThree Result | ' . ArrayMerge::getNiceFileSize(memory_get_usage(true)) . PHP_EOL;
$timeFour = microtime(true);
$d        = ArrayMerge::eachFour(10000);
echo 'count eachFour Result | ' . count($d) . PHP_EOL;
echo 'memory eachFour Result | ' . ArrayMerge::getNiceFileSize(memory_get_usage(true)) . PHP_EOL;
$timeFive = microtime(true);
echo PHP_EOL;
echo 'eachOne | ' . ($timeTwo - $timeOne) . PHP_EOL;
echo 'eachTwo | ' . ($timeThree - $timeTwo) . PHP_EOL;
echo 'eachThree | ' . ($timeFour - $timeThree) . PHP_EOL;
echo 'eachFour | ' . ($timeFive - $timeFour) . PHP_EOL;
echo PHP_EOL;

?>
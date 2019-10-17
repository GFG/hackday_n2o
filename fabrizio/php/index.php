<?php
require __DIR__ . '/vendor/autoload.php';

ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL);

use Aspera\Spreadsheet\XLSX\Reader;
use Aspera\Spreadsheet\XLSX\SharedStringsConfiguration;
use App\Connection;

$options = array(
    'TempDir'                    => '/tmp/',
    'SkipEmptyCells'             => false,
    'ReturnDateTimeObjects'      => true,
    'SharedStringsConfiguration' => new SharedStringsConfiguration(),
    'CustomFormats'              => array(20 => 'hh:mm')
);

try {
	define("COLUMN_SELLER_SKU", 0);
	define("COLUMN_SHOP_SKU", 1);
	define("COLUMN_PRICE", 2);
	define("COLUMN_SALE_PRICE", 3);
	define("COLUMN_SALE_START_DATE", 4);
	define("COLUMN_SALE_END_DATE", 5);
	define("COLUMN_NAME", 6);

    $start = microtime(true);
    $reader = new Reader($options);
	$reader->open('../static/PriceUpdate_30k.xlsx');

    foreach ($reader as $key => $row) {
        $list[] = [
			"COLUMN_SELLER_SKU" => $row[COLUMN_SELLER_SKU],
			"COLUMN_SHOP_SKU" => $row[COLUMN_SHOP_SKU],
			"COLUMN_PRICE" => $row[COLUMN_PRICE],
			"COLUMN_SALE_PRICE" => $row[COLUMN_SALE_PRICE],
			"COLUMN_SALE_START_DATE" => $row[COLUMN_SALE_START_DATE],
			"COLUMN_SALE_END_DATE" => $row[COLUMN_SALE_END_DATE],
			"COLUMN_NAME" => $row[COLUMN_NAME],
		];
    }

    $reader->close();
    $end = microtime(true);

    printf("\n\nNumber of products parsed: %.2f\n\n", ($end - $start)*1000);
}
catch (Exception $e) {
    die($e->getMessage());
}

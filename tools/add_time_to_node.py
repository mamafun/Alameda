import json
from data_processor import DataProcessor


def main(added_time_sec, input_file, output_file):

    starttime = DataProcessor.convert_time('2019-01-10T17:00:00Z')
    measurementname = 'node_prediction'
    with open(input_file) as inf:
        in_data = json.load(inf)

    for node in in_data['node_predictions']:
        for metric in node['predicted_raw_data']:
            idx = 0
            for sample in metric['data']:
                if 'CPU_USAGE_SECONDS_PERCENTAGE' == metric['metric_type']:
                    print(measurementname + ",node=" + node['name'] + ",type=cpu" + " value=" + sample['num_value'] + " " + str(starttime + idx*90) + "000000000")
                else:
                    print(measurementname + ",node=" + node['name'] + ",type=memory" + " value=" + sample['num_value'] + " " + str(starttime + idx*90) + "000000000")
                idx = idx + 1

#                time_in_sec = DataProcessor.convert_time(sample['time'])
#                sample['time'] = DataProcessor.convert_time(
#                    time_in_sec + added_time_sec)

    with open(output_file, 'w') as f:
        f.write(json.dumps(in_data))


if __name__ == '__main__':
    added_time_sec = 60
    input_file = 'node_predictions_before'
    output_file = input_file + '_time_changed'
    main(added_time_sec, input_file, output_file)

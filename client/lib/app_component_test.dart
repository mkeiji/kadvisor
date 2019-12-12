@TestOn('browser')
/* 
NOTE: chrome must be installed in the environment
run with: pub run build_runner test --fail-on-severe -- -p chrome
*/

import 'package:angular_app/app_component.dart';
import 'package:test/test.dart';

void main() {
    var component = AppComponent();

    setUp(() async {
        component.title = "test";
    });

    test('default', () {
        expect(component.title, "test");
    });
}

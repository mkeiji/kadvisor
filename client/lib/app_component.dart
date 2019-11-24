import 'package:angular/angular.dart';

import 'app-component.module.dart';

@Component(
    selector: 'my-app',
    styleUrls: ['app_component.css'],
    templateUrl: 'app_component.html',
    directives: [AppComponentDirectives],
)
class AppComponent {
    var title = 'Angular';
}
